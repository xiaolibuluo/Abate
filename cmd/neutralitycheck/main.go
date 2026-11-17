/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command neutralitycheck runs the eight go/ast zero-exemption invariants over
// every .go path registered in ci/e2e_sources.txt.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"abate/internal/e2e/neutrality"
)

const (
	exitOK        = 0
	exitFail      = 1
	exitUsage     = 2
	outputFileMod = 0o600
)

// neutralityOutput is the deterministic neutrality.json document.
type neutralityOutput struct {
	SchemaVersion   int                  `json:"schemaVersion"`
	RegisteredFiles int                  `json:"registeredFiles"`
	ProcessedFiles  int                  `json:"processedFiles"`
	Findings        []neutrality.Finding `json:"findings"`
	Passed          bool                 `json:"passed"`
}

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

// run parses the source list, checks every registered file and writes the
// result to -out. ProcessedFiles must equal RegisteredFiles to pass.
func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("neutralitycheck", flag.ContinueOnError)
	flags.SetOutput(stderr)
	sourcesPath := flags.String("sources", "ci/e2e_sources.txt",
		"path to ci/e2e_sources.txt")
	outPath := flags.String("out", "neutrality.json",
		"path to neutrality.json")
	if err := flags.Parse(args[1:]); err != nil {
		return exitUsage
	}
	paths, err := readPaths(*sourcesPath)
	if err != nil {
		emit(stderr, "neutralitycheck: %v\n", err)
		return exitFail
	}
	result, checkErr := neutrality.NewAnalyzer(".").Check(paths)
	doc := neutralityOutput{
		SchemaVersion: 1, RegisteredFiles: result.RegisteredFiles,
		ProcessedFiles: result.ProcessedFiles,
		Findings:       result.Findings,
		Passed: checkErr == nil && result.ProcessedFiles ==
			result.RegisteredFiles && len(result.Findings) == 0,
	}
	if err = writeDocument(*outPath, doc); err != nil {
		emit(stderr, "neutralitycheck: %v\n", err)
		return exitFail
	}
	for _, finding := range result.Findings {
		emit(stdout, "%s:%d %s %s\n", finding.File, finding.Line,
			finding.Rule, finding.Snippet)
	}
	if checkErr != nil {
		emit(stderr, "neutralitycheck: analysis failed: %v\n", checkErr)
		return exitFail
	}
	if !doc.Passed {
		emit(stderr,
			"neutralitycheck: FAIL registered=%d processed=%d findings=%d\n",
			result.RegisteredFiles, result.ProcessedFiles,
			len(result.Findings))
		return exitFail
	}
	emit(stdout, "NEUTRALITY_OK registered=%d processed=%d findings=0\n",
		result.RegisteredFiles, result.ProcessedFiles)
	return exitOK
}

// readPaths decodes the newline-separated registered source list.
func readPaths(path string) ([]string, error) {
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	out := make([]string, 0)
	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out, nil
}

// writeDocument encodes and stores neutrality.json.
func writeDocument(path string, doc neutralityOutput) error {
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return fmt.Errorf("encode neutrality: %w", err)
	}
	if err = os.WriteFile(path, append(data, '\n'), outputFileMod); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

// emit writes one line with an explicitly discarded write result.
func emit(w io.Writer, format string, args ...any) {
	_, _ = io.WriteString(w, fmt.Sprintf(format, args...))
}
