/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package targetgen

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

// Fingerprint returns the deterministic SHA-256 digest of the generated
// file set: file names are hashed in ascending order together with their
// bytes, so two generations agree on the fingerprint if and only if every
// file agrees byte for byte.
func (t Target) Fingerprint() [32]byte {
	names := make([]string, 0, len(t.Files))
	for name := range t.Files {
		names = append(names, name)
	}
	slices.Sort(names)
	h := sha256.New()
	for _, name := range names {
		h.Write([]byte(name))
		h.Write([]byte{0})
		h.Write(t.Files[name])
		h.Write([]byte{0})
	}
	var sum [32]byte
	copy(sum[:], h.Sum(nil))
	return sum
}

// ByteSize returns the total byte count of every generated file.
func (t Target) ByteSize() int {
	total := 0
	for _, data := range t.Files {
		total += len(data)
	}
	return total
}

// Write materializes the target project into dir atomically: the files land
// in a sibling temporary directory first, the optional self-check runs the
// real language toolchain against that staging tree, and only a fully green
// project is renamed into place. A failed write never leaves a partial
// target directory behind. dir must not exist yet, so a generation can
// never silently overwrite user sources.
func Write(dir string, target Target, selfCheck bool) error {
	if _, err := os.Stat(dir); err == nil {
		return fmt.Errorf("%w: %s", ErrDirExists, dir)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("targetgen: stat %s: %v", dir, err)
	}
	parent := filepath.Dir(dir)
	base := filepath.Base(dir)
	staging, err := os.MkdirTemp(parent, "."+base+".targetgen-")
	if err != nil {
		return fmt.Errorf("targetgen: staging directory: %v", err)
	}
	defer func() { _ = os.RemoveAll(staging) }()
	names := make([]string, 0, len(target.Files))
	for name := range target.Files {
		names = append(names, name)
	}
	slices.Sort(names)
	for _, name := range names {
		path := filepath.Join(staging, filepath.FromSlash(name))
		if err := os.WriteFile(path, target.Files[name], 0o644); err != nil {
			return fmt.Errorf("targetgen: write %s: %v", name, err)
		}
	}
	if selfCheck {
		if output, err := runSelfCheck(staging, target.Spec.Lang); err != nil {
			return fmt.Errorf("%w: %v: %s", ErrSelfCheck, err, tailOutput(output))
		}
	}
	if err := os.Rename(staging, dir); err != nil {
		return fmt.Errorf("targetgen: move into place %s: %v", dir, err)
	}
	return nil
}

// runSelfCheck executes the real language toolchain against the staged
// project. A red project is a generator bug and is reported through
// ErrSelfCheck instead of being delivered; a missing toolchain is reported
// through ErrToolchainFound so the caller can distinguish environment gaps
// from generator bugs. Build artifacts of the compiled languages live in a
// dedicated scratch directory removed afterwards, so the delivered project
// keeps exactly the generated source files.
func runSelfCheck(dir, lang string) (string, error) {
	scratch, err := os.MkdirTemp("", "abate-targetgen-check-")
	if err != nil {
		return "", fmt.Errorf("targetgen: self-check workspace: %v", err)
	}
	defer func() { _ = os.RemoveAll(scratch) }()
	steps, err := selfCheckSteps(dir, scratch, lang)
	if err != nil {
		return "", err
	}
	for _, step := range steps {
		if _, err := exec.LookPath(step.tool); err != nil {
			return "", fmt.Errorf("%w: %s", ErrToolchainFound, step.tool)
		}
		cmd := exec.Command(step.tool, step.args...)
		cmd.Dir = dir
		var buf strings.Builder
		cmd.Stdout = &buf
		cmd.Stderr = &buf
		if err := cmd.Run(); err != nil {
			return buf.String(), fmt.Errorf("%s %s: %v", step.tool, strings.Join(step.args, " "), err)
		}
	}
	return "", nil
}

// checkStep is one toolchain invocation of a self-check sequence: compiled
// languages need a build step before the run step, interpreted languages
// run their single file directly.
type checkStep struct {
	tool string
	args []string
}

// selfCheckSteps returns the ordered toolchain invocations proving one
// staged project: the project compiles (where the language requires it) and
// its in-file tests all pass. Every command runs offline; compiled
// artifacts land in scratch, never in the delivered tree.
func selfCheckSteps(dir, scratch, lang string) ([]checkStep, error) {
	switch lang {
	case LangGo:
		return []checkStep{{tool: "go", args: []string{"test", "-count=1", "./..."}}}, nil
	case LangPython:
		return []checkStep{{tool: "python3", args: []string{"target.py"}}}, nil
	case LangJS:
		return []checkStep{{tool: "node", args: []string{"target.js"}}}, nil
	case LangC:
		binary := filepath.Join(scratch, "c-target")
		return []checkStep{
			{tool: "gcc", args: []string{"-O1", "-o", binary, "target.c"}},
			{tool: binary},
		}, nil
	case LangCSharp:
		return csharpSteps(dir, scratch)
	case LangRust:
		binary := filepath.Join(scratch, "rs-target")
		return []checkStep{
			{tool: "rustc", args: []string{"-O", "-o", binary, "target.rs"}},
			{tool: binary},
		}, nil
	case LangKotlin:
		jar := filepath.Join(scratch, "kt-target.jar")
		return []checkStep{
			{tool: "kotlinc", args: []string{"target.kt", "-include-runtime", "-d", jar}},
			{tool: "java", args: []string{"-jar", jar}},
		}, nil
	}
	// Unreachable: Validate admits only registered languages.
	panic("targetgen: unknown language " + lang)
}

// csharpSteps prepares a minimal throwaway project around the generated
// file and returns the run step. File-based execution (`dotnet run file.cs`)
// only exists on .NET 10+, so the check stages a scratch csproj whose
// target framework matches the installed SDK instead; the scratch tree is
// removed with the self-check workspace and never enters the delivery.
func csharpSteps(dir, scratch string) ([]checkStep, error) {
	source, err := os.ReadFile(filepath.Join(dir, "target.cs"))
	if err != nil {
		return nil, fmt.Errorf("targetgen: stage C# self-check: %v", err)
	}
	proj := filepath.Join(scratch, "cs")
	if err := os.MkdirAll(proj, 0o755); err != nil {
		return nil, fmt.Errorf("targetgen: stage C# self-check: %v", err)
	}
	if err := os.WriteFile(filepath.Join(proj, "target.cs"), source, 0o644); err != nil {
		return nil, fmt.Errorf("targetgen: stage C# self-check: %v", err)
	}
	tfm, err := dotnetTargetFramework()
	if err != nil {
		return nil, err
	}
	csproj := `<Project Sdk="Microsoft.NET.Sdk"><PropertyGroup>` +
		`<OutputType>Exe</OutputType>` +
		`<TargetFramework>` + tfm + `</TargetFramework>` +
		`<Nullable>disable</Nullable><ImplicitUsings>disable</ImplicitUsings>` +
		`</PropertyGroup></Project>`
	if err := os.WriteFile(filepath.Join(proj, "targetgen-check.csproj"), []byte(csproj), 0o644); err != nil {
		return nil, fmt.Errorf("targetgen: stage C# self-check: %v", err)
	}
	return []checkStep{{tool: "dotnet", args: []string{"run", "--project", proj}}}, nil
}

// dotnetTargetFramework maps the installed dotnet SDK version onto the
// matching target framework moniker ("8.0.428" -> "net8.0"), so the scratch
// project restores against the locally installed framework reference packs
// without any network access.
func dotnetTargetFramework() (string, error) {
	out, err := exec.Command("dotnet", "--version").Output()
	if err != nil {
		return "", fmt.Errorf("%w: dotnet", ErrToolchainFound)
	}
	parts := strings.SplitN(strings.TrimSpace(string(out)), ".", 3)
	if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
		return "", fmt.Errorf("targetgen: cannot parse dotnet version %q", strings.TrimSpace(string(out)))
	}
	return "net" + parts[0] + "." + parts[1], nil
}

// tailOutput truncates self-check output so an error message stays readable
// while still carrying the failing test diagnostics.
func tailOutput(output string) string {
	text := strings.TrimSpace(output)
	if len(text) > 512 {
		return text[:512]
	}
	return text
}
