/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"encoding/json"
	"fmt"

	"abate/kernel/mutation"
	"abate/lang/rule"
)

// BuildOTel renders findings and a mutation report into a deterministic,
// air-gapped OTLP-style JSON payload. It performs no export: there is no network
// field, no SDK dependency and no wall-clock timestamp, and time is expressed
// solely through integer virtualTime values derived from the report.
//
// 参考 OTLP/JSON 的 span / metric / log 三流结构，改进点=完全离线序列化，
// 虚拟时间替代真实时间戳，导出动作退化为返回字节切片。
// The returned bytes are a pure function of the two inputs.
func BuildOTel(findings []rule.Finding, rep mutation.Report) ([]byte, error) {
	doc := otelDocument{
		ResourceSpans: []otelResourceSpans{buildOTelResourceSpans(rep)},
		Metrics:       buildOTelMetrics(findings, rep),
		Logs:          buildOTelLogs(findings, rep),
	}
	body, err := json.Marshal(doc)
	if err != nil {
		return nil, fmt.Errorf("report: marshal otel payload: %w", err)
	}
	return body, nil
}
