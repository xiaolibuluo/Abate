/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"encoding/json"

	"abate/kernel/mutation"
	"abate/lang/rule"
)

// BuildArgs is the serializable payload passed to Host.Call for EntryBuild. It
// carries the static findings and the mutation report that a ReporterPlugin
// renders into an artifact (for example a SARIF document). Encoding the payload
// once, in a deterministic JSON form, keeps the host.Call signature uniform
// across every plugin type while remaining byte-stable for replay.
type BuildArgs struct {
	// Findings are the static analysis matches to be reported.
	Findings []rule.Finding

	// Report is the deterministic mutation report to be rendered.
	Report mutation.Report
}

// EncodeBuildArgs serializes a BuildArgs payload into a compact, deterministic
// JSON document suitable for transport through Host.Call args.
func EncodeBuildArgs(a BuildArgs) ([]byte, error) {
	return json.Marshal(a)
}

// DecodeBuildArgs parses a previously encoded BuildArgs payload. A malformed or
// non-JSON payload is rejected so callers cannot invoke a reporter with garbage.
func DecodeBuildArgs(data []byte, out *BuildArgs) error {
	return json.Unmarshal(data, out)
}
