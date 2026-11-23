/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"abate/kernel/mutation"
	"abate/lang/cmutate"
	"abate/lang/csmutate"
	"abate/lang/gomutate"
	"abate/lang/jsmutate"
	"abate/lang/ktmutate"
	"abate/lang/mutcore"
	"abate/lang/pymutate"
	"abate/lang/rsmutate"
)

// ErrUnsupportedLanguage reports a mutate target whose language backend is
// not implemented or cannot be detected from the workspace content.
var ErrUnsupportedLanguage = errors.New("cli: unsupported mutation language")

// mutationBackend is the language-agnostic view of one real-workspace
// mutation session the CLI consumes. gomutate.Backend satisfies it
// directly; mutcore-based backends are adapted through sessionBackend.
type mutationBackend interface {
	Target() mutation.Target
	Suite() mutation.Suite
	Close() error
	Root() string
	ModulePath() string
	Language() string
	TokenCount() int
	TestCount() int
}

// sessionBackend adapts a mutcore session to the CLI backend view. The
// module path of a non-Go workspace is its root directory: the concept of
// a module manifest is Go-specific, and the wire view keeps the field for
// report continuity rather than inventing per-language manifests.
type sessionBackend struct {
	session *mutcore.Session
}

// Target returns the mutation.Target view of the session.
func (b sessionBackend) Target() mutation.Target { return b.session.Target() }

// Suite returns the mutation.Suite view of the session.
func (b sessionBackend) Suite() mutation.Suite { return b.session.Suite() }

// Close releases the staged mutant workspaces of the session.
func (b sessionBackend) Close() error { return b.session.Close() }

// Root exposes the absolute workspace root the session mutates.
func (b sessionBackend) Root() string { return b.session.Root() }

// ModulePath exposes the workspace root as the module identity of non-Go
// backends.
func (b sessionBackend) ModulePath() string { return b.session.Root() }

// Language exposes the canonical tag of the session driver.
func (b sessionBackend) Language() string { return b.session.Language() }

// TokenCount reports how many mutable tokens the scanner discovered.
func (b sessionBackend) TokenCount() int { return b.session.TokenCount() }

// TestCount reports how many tests the runner discovered.
func (b sessionBackend) TestCount() int { return b.session.TestCount() }

// goBackend wraps gomutate.Backend with the Language method the generic
// view requires, without touching the Go backend itself.
type goBackend struct {
	backend *gomutate.Backend
}

// Target returns the mutation.Target view of the Go session.
func (b goBackend) Target() mutation.Target { return b.backend.Target() }

// Suite returns the mutation.Suite view of the Go session.
func (b goBackend) Suite() mutation.Suite { return b.backend.Suite() }

// Close releases the staged mutant workspaces of the Go session.
func (b goBackend) Close() error { return b.backend.Close() }

// Root exposes the absolute workspace root the Go session mutates.
func (b goBackend) Root() string { return b.backend.Root() }

// ModulePath exposes the resolved module path of the Go workspace.
func (b goBackend) ModulePath() string { return b.backend.ModulePath() }

// Language returns the canonical Go tag.
func (b goBackend) Language() string { return "go" }

// TokenCount reports how many mutable tokens the Go scanner discovered.
func (b goBackend) TokenCount() int { return b.backend.TokenCount() }

// TestCount reports how many tests the Go toolchain discovered.
func (b goBackend) TestCount() int { return b.backend.TestCount() }

// newMutationBackend builds the real backend for one language tag over the
// target workspace. An empty tag triggers content-based detection, so an
// explicit --lang never surprises a workspace that already knows what it
// is. Every constructor failure propagates unchanged: a missing toolchain
// stays ErrToolchainMissing, an empty suite stays ErrNoTests, and an
// unknown tag is rejected with ErrUnsupportedLanguage.
func newMutationBackend(ctx context.Context, lang, path string, native bool) (mutationBackend, error) {
	resolved := lang
	if resolved == "" {
		detected, err := detectMutationLanguage(path)
		if err != nil {
			return nil, err
		}
		resolved = detected
	}
	switch resolved {
	case "go":
		if native {
			backend, err := gomutate.NewNative(ctx, path)
			if err != nil {
				return nil, err
			}
			return goBackend{backend: backend}, nil
		}
		backend, err := gomutate.New(ctx, path)
		if err != nil {
			return nil, err
		}
		return goBackend{backend: backend}, nil
	case "py":
		if native {
			session, err := pymutate.NewNative(ctx, path)
			if err != nil {
				return nil, err
			}
			return sessionBackend{session: session}, nil
		}
		session, err := pymutate.New(ctx, path)
		if err != nil {
			return nil, err
		}
		return sessionBackend{session: session}, nil
	case "js":
		if native {
			session, err := jsmutate.NewNative(ctx, path)
			if err != nil {
				return nil, err
			}
			return sessionBackend{session: session}, nil
		}
		session, err := jsmutate.New(ctx, path)
		if err != nil {
			return nil, err
		}
		return sessionBackend{session: session}, nil
	case "c":
		if native {
			session, err := cmutate.NewNative(ctx, path)
			if err != nil {
				return nil, err
			}
			return sessionBackend{session: session}, nil
		}
		session, err := cmutate.New(ctx, path)
		if err != nil {
			return nil, err
		}
		return sessionBackend{session: session}, nil
	case "rs":
		if native {
			session, err := rsmutate.NewNative(ctx, path)
			if err != nil {
				return nil, err
			}
			return sessionBackend{session: session}, nil
		}
		session, err := rsmutate.New(ctx, path)
		if err != nil {
			return nil, err
		}
		return sessionBackend{session: session}, nil
	case "cs":
		if native {
			session, err := csmutate.NewNative(ctx, path)
			if err != nil {
				return nil, err
			}
			return sessionBackend{session: session}, nil
		}
		session, err := csmutate.New(ctx, path)
		if err != nil {
			return nil, err
		}
		return sessionBackend{session: session}, nil
	case "kt":
		if native {
			session, err := ktmutate.NewNative(ctx, path)
			if err != nil {
				return nil, err
			}
			return sessionBackend{session: session}, nil
		}
		session, err := ktmutate.New(ctx, path)
		if err != nil {
			return nil, err
		}
		return sessionBackend{session: session}, nil
	}
	return nil, fmt.Errorf("%w: %q", ErrUnsupportedLanguage, resolved)
}

// detectMutationLanguage infers the workspace language from its content:
// a go.mod anchors a Go module, any regular .py file marks a Python
// workspace. The check is deterministic (lexicographic directory order,
// first match wins) so two runs over the same tree never disagree.
func detectMutationLanguage(path string) (string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("%w: %v", mutcore.ErrInvalidRoot, err)
	}
	if info, err := os.Stat(filepath.Join(path, "go.mod")); err == nil && !info.IsDir() {
		return "go", nil
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.EqualFold(filepath.Ext(entry.Name()), ".py") {
			return "py", nil
		}
		if strings.EqualFold(filepath.Ext(entry.Name()), ".js") {
			return "js", nil
		}
		if strings.EqualFold(filepath.Ext(entry.Name()), ".c") {
			return "c", nil
		}
		if strings.EqualFold(filepath.Ext(entry.Name()), ".rs") {
			return "rs", nil
		}
		if strings.EqualFold(filepath.Ext(entry.Name()), ".cs") {
			return "cs", nil
		}
		if strings.EqualFold(filepath.Ext(entry.Name()), ".kt") {
			return "kt", nil
		}
	}
	return "", fmt.Errorf("%w: cannot detect language of %s",
		ErrUnsupportedLanguage, path)
}
