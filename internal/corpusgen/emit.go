/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// errGenerate reports a generator-side failure that must never be converted
// into a silently shorter corpus.
var errGenerate = errors.New("corpusgen: generation failed")

const (
	generatedFilePerm  = 0o600
	generatedDirPerm   = 0o750
	maxSemanticBlocks  = 10000
	largeFileThreshold = 8192
	convergenceSpan    = 3400
)

// longNameWords are ordinary descriptive camel-case words. Long identifiers
// are legal code tokens used only for final byte-band convergence; they keep
// the AST node count of 160KB files low enough for the runner budget.
var longNameWords = []string{
	"Probe", "Relay", "Channel", "Select", "Dispatch", "Accumulate",
	"Transform", "Describe", "Classify", "Encode", "Fold", "Scan",
	"Store", "Value", "Result", "Sequence", "Deterministic", "Buffer",
	"Cursor", "Matrix", "Vector", "Sampler", "Reader", "Writer",
	"Mapper", "Reducer", "Filter", "Gateway", "Bridge", "Adapter",
	"Handler", "Emitter", "Collector", "Resolver", "Validator",
	"Generator", "Iterator", "Matcher", "Visitor", "Scheduler",
	"Registry", "Factory", "Builder", "Parser", "Decoder",
	"Comparator", "Aggregator", "Calculator", "Converter",
}

// longIdentifier builds a deterministic, unique identifier of the requested
// length from ordinary code words plus the file index and block sequence.
func longIdentifier(index, seq, want int, style IdentifierStyle) string {
	snake := style == StyleSnake
	var out strings.Builder
	for out.Len() < want {
		word := longNameWords[(index*7+seq*11+out.Len()/8)%len(longNameWords)]
		if snake {
			out.WriteString(strings.ToLower(word))
			out.WriteByte('_')
		} else {
			out.WriteString(word)
		}
	}
	out.WriteString(itoa(index))
	out.WriteByte('x')
	out.WriteString(itoa(seq))
	return out.String()
}

// convergenceBlock renders one compact function whose identifier carries the
// byte band. The identifier is a real token, not a comment or whitespace run.
func convergenceBlock(language string, index, seq, span int) string {
	style := StyleCamel
	if language == "py" || language == "c" || language == "rs" {
		style = StyleSnake
	}
	name := longIdentifier(index, seq, span, style)
	template := map[string]string{
		"go": "func @@L@@() int {\n\treturn 1\n}\n",
		"py": "def @@L@@():\n    return 1\n",
		"js": "function @@L@@() { return 1; }\n",
		"c":  "int @@L@@(void) { return 1; }\n",
		"cs": "    public static int @@L@@() { return 1; }\n",
		"rs": "pub fn @@L@@() -> i32 { 1 }\n",
		"kt": "fun @@L@@(): Int = 1\n",
	}[language]
	if template == "" {
		return ""
	}
	return strings.ReplaceAll(template, "@@L@@", name)
}

// renderFile builds one high-fidelity file by appending real semantic blocks
// until the byte target is reached, then asserts the capture anchor.
func renderFile(recipe Recipe, mode AnchorMode, target int, rng *LabeledRNG,
	index int) ([]byte, error) {
	header := renderPlaceholders(recipe.Header, index, 0)
	if HeaderHasToken(header) {
		return nil, fmt.Errorf("%w: %s header contains a mutation token",
			errGenerate, recipe.Language)
	}
	var out strings.Builder
	out.WriteString(header)
	start := out.Len()
	anchor := recipe.AnchorFull
	if mode == AnchorClean {
		anchor = recipe.AnchorClean
	}
	out.WriteString(anchor)
	end := out.Len()
	// Always append at least one semantic block: an anchor-only file would be
	// byte-identical across indices and break the SHA-uniqueness contract.
	for seq := 0; seq < 1 || out.Len() < target; seq++ {
		if seq >= maxSemanticBlocks {
			return nil, fmt.Errorf("%w: %s convergence exhausted",
				errGenerate, recipe.Language)
		}
		if mode == AnchorClean {
			out.WriteString(recipe.RenderCleanBlock(rng, index, seq, mode))
			continue
		}
		if target > largeFileThreshold && seq >= 3 {
			out.WriteString(convergenceBlock(recipe.Language, index, seq,
				convergenceSpan))
			continue
		}
		out.WriteString(recipe.RenderBlock(rng, index, seq))
	}
	out.WriteString(recipe.Footer)
	data := []byte(out.String())
	bounds := CaptureAnchor{Language: recipe.Language, Clean: mode == AnchorClean,
		Bounds: AnchorBounds{Start: start, End: end}}
	if err := bounds.Preflight(data); err != nil {
		return nil, err
	}
	return data, nil
}

// assertEmptyRoot refuses to mix a new generation with pre-existing fixtures.
func assertEmptyRoot(root string) error {
	info, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("%w: stat %s: %w", errGenerate, root, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%w: root %s is not a directory", errGenerate, root)
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return fmt.Errorf("%w: read %s: %w", errGenerate, root, err)
	}
	if len(entries) != 0 {
		return fmt.Errorf("%w: root %s is not empty; remove it before generating",
			errGenerate, root)
	}
	return nil
}

// generationContext bundles the immutable inputs of one generation run.
type generationContext struct {
	ref     PlanRef
	bands   SizeBands
	opts    GenOptions
	recipes map[string]Recipe
	quotas  map[string]int
	seen    map[string]string
	stats   GenerateStats
}

// newGenerationContext validates options and prepares the output root.
func newGenerationContext(opts GenOptions) (generationContext, error) {
	ctx := generationContext{opts: opts,
		recipes: map[string]Recipe{}, seen: map[string]string{},
		stats: GenerateStats{PerLanguage: map[string]int{}, PerStratum: map[string]int{}}}
	ref, err := LoadPlanRef(opts.PlanPath)
	if err != nil {
		return ctx, err
	}
	if opts.Strata != ref.StrataCount || opts.Count < len(ref.Languages) {
		return ctx, fmt.Errorf("%w: invalid count/strata", errGenerate)
	}
	bands, err := NewSizeBands(ref.MinFileBytes, ref.MaxFileBytes, ref.StrataCount)
	if err != nil {
		return ctx, err
	}
	if err = assertEmptyRoot(opts.Root); err != nil {
		return ctx, err
	}
	if err = os.MkdirAll(opts.Root, generatedDirPerm); err != nil {
		return ctx, fmt.Errorf("%w: mkdir %s: %w", errGenerate, opts.Root, err)
	}
	for _, recipe := range AllRecipes() {
		ctx.recipes[recipe.Language] = recipe
	}
	ctx.ref = ref
	ctx.bands = bands
	ctx.quotas = languageQuotas(opts.Count, ref.Languages)
	return ctx, nil
}

// writeOneFile renders, validates and persists one corpus source file.
func (c *generationContext) writeOneFile(language string, recipe Recipe,
	stratum, index int, relPath string) error {
	rng := NewLabeledRNG(c.ref.Seed,
		fmt.Sprintf("%s:%d:%d", language, stratum, index))
	mode := AnchorFull
	if stratum == 0 && index == 0 {
		mode = AnchorClean
	}
	data, err := renderFile(recipe, mode, c.bands.Target(stratum, index),
		&rng, index)
	if err != nil {
		return err
	}
	if actual := c.bands.StratumOf(len(data)); actual != stratum {
		return fmt.Errorf("%w: %s band %d != %d", errGenerate, relPath, actual, stratum)
	}
	sum := sha256.Sum256(data)
	digest := hex.EncodeToString(sum[:])
	if previous, duplicated := c.seen[digest]; duplicated {
		return fmt.Errorf("%w: duplicate sha %s between %s and %s",
			errGenerate, digest, previous, relPath)
	}
	c.seen[digest] = relPath
	if err = os.WriteFile(filepath.Join(c.opts.Root, relPath), data,
		generatedFilePerm); err != nil {
		return fmt.Errorf("%w: write %s: %w", errGenerate, relPath, err)
	}
	c.stats.Total++
	c.stats.Bytes += int64(len(data))
	c.stats.PerLanguage[language]++
	c.stats.PerStratum[fmt.Sprintf("%s:s%d", language, stratum)]++
	return nil
}

// writeStratum writes one language's one byte band.
func (c *generationContext) writeStratum(language string, recipe Recipe,
	stratum int) error {
	dir := filepath.Join(c.opts.Root, language, fmt.Sprintf("s%d", stratum))
	if err := os.MkdirAll(dir, generatedDirPerm); err != nil {
		return fmt.Errorf("%w: mkdir %s: %w", errGenerate, dir, err)
	}
	quota := c.bands.Quota(c.quotas[language], stratum)
	if quota < 1 {
		return fmt.Errorf("%w: empty quota %s/s%d", errGenerate, language, stratum)
	}
	for index, relPath := range c.bands.Names(language, recipe.Ext, stratum, quota) {
		if err := c.writeOneFile(language, recipe, stratum, index, relPath); err != nil {
			return err
		}
	}
	return nil
}

// Generate writes the complete self-researched corpus and its QA fixtures.
func Generate(opts GenOptions) (GenerateStats, error) {
	ctx, err := newGenerationContext(opts)
	if err != nil {
		return ctx.stats, err
	}
	for _, language := range ctx.ref.Languages {
		recipe, found := ctx.recipes[language]
		if !found {
			return ctx.stats, fmt.Errorf("%w: no recipe for %s", errGenerate, language)
		}
		langDir := filepath.Join(opts.Root, language)
		if err = os.MkdirAll(langDir, generatedDirPerm); err != nil {
			return ctx.stats, fmt.Errorf("%w: mkdir %s: %w", errGenerate, langDir, err)
		}
		for stratum := 0; stratum < ctx.ref.StrataCount; stratum++ {
			if err = ctx.writeStratum(language, recipe, stratum); err != nil {
				return ctx.stats, err
			}
		}
	}
	dedicationPath := filepath.Join(opts.Root, DedicationRelPath)
	if err = os.WriteFile(dedicationPath, []byte(DedicationText),
		generatedFilePerm); err != nil {
		return ctx.stats, fmt.Errorf("%w: write dedication: %w", errGenerate, err)
	}
	if opts.BrokenRoot != "" {
		ctx.stats.Broken, err = GenerateBroken(opts.BrokenRoot, opts.Root, ctx.ref)
		if err != nil {
			return ctx.stats, err
		}
	}
	return ctx.stats, nil
}
