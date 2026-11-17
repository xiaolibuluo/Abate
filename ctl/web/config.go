/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package web

import (
	"flag"
	"io"
	"strings"
)

// Defaults for the web command. They are frozen so repeated invocations with
// identical arguments produce identical bytes.
const (
	// DefaultListen binds the loopback interface only, matching the safe-by-default
	// posture of the control plane (PLAN-07 CTL-WEB section 4).
	DefaultListen = "127.0.0.1:8430"
	// DefaultRestBase is the empty string: the internal REST server is embedded.
	DefaultRestBase = ""
)

// Exit codes mirror the CLI contract so "abate web" composes with other commands.
const (
	exitWebOK         = 0
	exitWebInternal   = 2
	exitWebUsageError = 3
)

// Config holds the parsed web command configuration.
type Config struct {
	// Listen is the host:port to bind.
	Listen string
	// Token is the bearer token required for API access. It is mandatory for a
	// non-loopback bind address; the embedded REST server rejects such a start
	// without a token, so the service can never come up exposed and unauthenticated.
	Token string
	// RestBase is an external REST base URL. When empty, the internal REST server
	// is embedded; otherwise /api is reverse-proxied to this base.
	RestBase string
	// StaticDir overrides the embedded assets with a directory during development.
	StaticDir string
}

// ParseConfig parses web flags with deterministic, ContinueOnError semantics.
//
// A parse or validation failure writes the help text to stderr and returns the
// usage exit code, exactly like the other CLI commands.
func ParseConfig(args []string, stderr io.Writer) (Config, int) {
	fs := flag.NewFlagSet("web", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	var cfg Config
	fs.StringVar(&cfg.Listen, "listen", DefaultListen, "host:port to bind")
	fs.StringVar(&cfg.Token, "token", "", "bearer token required for API access")
	fs.StringVar(&cfg.RestBase, "rest", DefaultRestBase, "external REST base URL; empty embeds the internal server")
	fs.StringVar(&cfg.StaticDir, "static", "", "directory overriding the embedded static assets")
	if err := fs.Parse(args); err != nil {
		_, _ = io.WriteString(stderr, "abate web: "+err.Error()+"\n")
		writeWebHelp(stderr)
		return Config{}, exitWebUsageError
	}
	if strings.TrimSpace(cfg.Listen) == "" {
		_, _ = io.WriteString(stderr, "abate web: empty listen address\n")
		writeWebHelp(stderr)
		return Config{}, exitWebUsageError
	}
	return cfg, exitWebOK
}

// helpRequested reports whether an argument slice asks for command help.
func helpRequested(args []string) bool {
	for _, a := range args {
		if a == "-h" || a == "--help" || a == "-help" {
			return true
		}
	}
	return false
}
