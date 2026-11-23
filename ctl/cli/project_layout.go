/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"os"
	"path/filepath"
	"strings"
)

// ProjectLayout describes the detected project type and whether it
// benefits from native sandbox mode. Native mode reduces per-mutant
// workspace overhead by linking heavy dependency directories (vendor,
// node_modules, target) once instead of copying them per mutant.
type ProjectLayout struct {
	// Language is the canonical language tag (go, py, js, c, rs, cs, kt).
	Language string
	// NativeCompatible reports whether the project layout supports
	// native sandbox mode: it has a recognizable dependency directory
	// that can be linked once rather than copied per mutant.
	NativeCompatible bool
	// Reason explains why the layout is or is not native-compatible,
	// used in CLI warnings when the user passes --native on an
	// incompatible project.
	Reason string
}

// detectProjectLayout scans the workspace root and determines the project
// type and native-mode compatibility. The detection is deterministic: it
// checks manifest files in lexicographic order and returns the first
// match, so two runs over the same tree never disagree.
func detectProjectLayout(path string) (ProjectLayout, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return ProjectLayout{}, err
	}

	// Go: go.mod anchors a module. vendor/ presence makes it
	// native-compatible because vendor is the heavy directory that
	// benefits from one-time linking.
	if info, err := os.Stat(filepath.Join(path, "go.mod")); err == nil && !info.IsDir() {
		hasVendor := false
		if info, err := os.Stat(filepath.Join(path, "vendor")); err == nil && info.IsDir() {
			hasVendor = true
		}
		reason := "go.mod found"
		if hasVendor {
			reason += " with vendor/ directory"
		}
		return ProjectLayout{
			Language:         "go",
			NativeCompatible: true,
			Reason:           reason,
		}, nil
	}

	// Python: requirements.txt, setup.py, or pyproject.toml mark a
	// Python project. site-packages or venv/ make it native-compatible.
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := strings.ToLower(entry.Name())
		if name == "requirements.txt" || name == "setup.py" || name == "pyproject.toml" {
			hasDeps := false
			for _, depDir := range []string{"site-packages", "venv", ".venv", "__pycache__"} {
				if info, err := os.Stat(filepath.Join(path, depDir)); err == nil && info.IsDir() {
					hasDeps = true
					break
				}
			}
			reason := "Python manifest found"
			if hasDeps {
				reason += " with dependency directories"
			}
			return ProjectLayout{
				Language:         "py",
				NativeCompatible: hasDeps,
				Reason:           reason,
			}, nil
		}
	}

	// JavaScript/Node: package.json marks a Node project. node_modules/
	// makes it native-compatible.
	if info, err := os.Stat(filepath.Join(path, "package.json")); err == nil && !info.IsDir() {
		hasNodeModules := false
		if info, err := os.Stat(filepath.Join(path, "node_modules")); err == nil && info.IsDir() {
			hasNodeModules = true
		}
		reason := "package.json found"
		if hasNodeModules {
			reason += " with node_modules/"
		}
		return ProjectLayout{
			Language:         "js",
			NativeCompatible: hasNodeModules,
			Reason:           reason,
		}, nil
	}

	// Rust: Cargo.toml marks a Rust project. target/ makes it
	// native-compatible.
	if info, err := os.Stat(filepath.Join(path, "Cargo.toml")); err == nil && !info.IsDir() {
		hasTarget := false
		if info, err := os.Stat(filepath.Join(path, "target")); err == nil && info.IsDir() {
			hasTarget = true
		}
		reason := "Cargo.toml found"
		if hasTarget {
			reason += " with target/"
		}
		return ProjectLayout{
			Language:         "rs",
			NativeCompatible: hasTarget,
			Reason:           reason,
		}, nil
	}

	// C# / .NET: .csproj or .sln marks a C# project. bin/ or obj/
	// makes it native-compatible.
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext == ".csproj" || ext == ".sln" {
			hasBuild := false
			for _, depDir := range []string{"bin", "obj"} {
				if info, err := os.Stat(filepath.Join(path, depDir)); err == nil && info.IsDir() {
					hasBuild = true
					break
				}
			}
			reason := ".NET project file found"
			if hasBuild {
				reason += " with bin/obj directories"
			}
			return ProjectLayout{
				Language:         "cs",
				NativeCompatible: hasBuild,
				Reason:           reason,
			}, nil
		}
	}

	// Kotlin: build.gradle.kts or settings.gradle.kts marks a Kotlin
	// project. .gradle/ or build/ makes it native-compatible.
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := strings.ToLower(entry.Name())
		if name == "build.gradle.kts" || name == "settings.gradle.kts" {
			hasBuild := false
			for _, depDir := range []string{"build", ".gradle"} {
				if info, err := os.Stat(filepath.Join(path, depDir)); err == nil && info.IsDir() {
					hasBuild = true
					break
				}
			}
			reason := "Kotlin build file found"
			if hasBuild {
				reason += " with build/.gradle directories"
			}
			return ProjectLayout{
				Language:         "kt",
				NativeCompatible: hasBuild,
				Reason:           reason,
			}, nil
		}
	}

	// C: Makefile or CMakeLists.txt marks a C project. build/ or
	// .objs/ makes it native-compatible.
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := strings.ToLower(entry.Name())
		if name == "makefile" || name == "cmakelists.txt" {
			hasBuild := false
			for _, depDir := range []string{"build", ".objs", "objs"} {
				if info, err := os.Stat(filepath.Join(path, depDir)); err == nil && info.IsDir() {
					hasBuild = true
					break
				}
			}
			reason := "C build file found"
			if hasBuild {
				reason += " with build directories"
			}
			return ProjectLayout{
				Language:         "c",
				NativeCompatible: hasBuild,
				Reason:           reason,
			}, nil
		}
	}

	// Fallback: check for source files to infer language. Native
	// compatibility is false because no dependency directory was
	// detected.
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		switch ext {
		case ".go":
			return ProjectLayout{
				Language:         "go",
				NativeCompatible: false,
				Reason:           "Go source files found but no go.mod (not a module)",
			}, nil
		case ".py":
			return ProjectLayout{
				Language:         "py",
				NativeCompatible: false,
				Reason:           "Python source files found but no manifest",
			}, nil
		case ".js":
			return ProjectLayout{
				Language:         "js",
				NativeCompatible: false,
				Reason:           "JavaScript source files found but no package.json",
			}, nil
		case ".c":
			return ProjectLayout{
				Language:         "c",
				NativeCompatible: false,
				Reason:           "C source files found but no Makefile/CMakeLists.txt",
			}, nil
		case ".rs":
			return ProjectLayout{
				Language:         "rs",
				NativeCompatible: false,
				Reason:           "Rust source files found but no Cargo.toml",
			}, nil
		case ".cs":
			return ProjectLayout{
				Language:         "cs",
				NativeCompatible: false,
				Reason:           "C# source files found but no .csproj/.sln",
			}, nil
		case ".kt":
			return ProjectLayout{
				Language:         "kt",
				NativeCompatible: false,
				Reason:           "Kotlin source files found but no build.gradle.kts",
			}, nil
		}
	}

	return ProjectLayout{
		Language:         "",
		NativeCompatible: false,
		Reason:           "no recognizable project layout",
	}, nil
}
