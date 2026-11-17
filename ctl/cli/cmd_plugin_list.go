/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"sort"

	"abate/plugin/signature"
)

// pluginListItem is one row of the deterministic registry listing.
type pluginListItem struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

// cmdPluginList prints every registered instance in id order. Sorting is
// mandatory: Go map iteration order is randomized and must never leak into CLI
// output, which CI compares byte for byte.
func cmdPluginList(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writePluginHelp(stdout)
		return exitOK
	}
	fs := flag.NewFlagSet("plugin list", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	if err := fs.Parse(args); err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin list: %v\n", err)
		return exitUsageError
	}
	if fs.NArg() != 0 {
		writePluginHelp(stderr)
		return exitUsageError
	}

	session, err := loadPluginSession()
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate plugin list: %v\n", err)
		return exitInternalError
	}
	snapshot := session.h.Registry().Snapshot()

	ids := make([]string, 0, len(snapshot))
	for id := range snapshot {
		ids = append(ids, string(id))
	}
	sort.Strings(ids)
	items := make([]pluginListItem, 0, len(ids))
	for _, id := range ids {
		inst := snapshot[signature.PluginID(id)]
		if inst == nil {
			continue
		}
		items = append(items, pluginListItem{
			ID:      string(inst.Manifest.ID),
			Type:    string(inst.Manifest.Type),
			Version: inst.Manifest.Version.String(),
		})
	}
	data, err := json.Marshal(items)
	if err != nil {
		return exitInternalError
	}
	if _, err := stdout.Write(append(data, '\n')); err != nil {
		return exitInternalError
	}
	return exitOK
}
