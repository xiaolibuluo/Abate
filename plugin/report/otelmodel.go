/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

// otelDocument is the air-gapped OTLP-style payload with three ordered streams.
// It contains no export destination and no wall-clock timestamp anywhere.
type otelDocument struct {
	ResourceSpans []otelResourceSpans `json:"resourceSpans"`
	Metrics       []otelMetric        `json:"metrics"`
	Logs          []otelLogRecord     `json:"logs"`
}

// otelResourceSpans wraps one resource and its deterministic span tree.
type otelResourceSpans struct {
	Resource otelResource `json:"resource"`
	Spans    []otelSpan   `json:"spans"`
}

// otelResource carries only deterministic identifiers of the producer.
type otelResource struct {
	Attributes []otelAttribute `json:"attributes"`
}

// otelSpan is one node of the deterministic HPOV phase tree.
type otelSpan struct {
	TraceID      string          `json:"traceId"`
	SpanID       string          `json:"spanId"`
	ParentSpanID string          `json:"parentSpanId,omitempty"`
	Name         string          `json:"name"`
	VirtualTime  string          `json:"virtualTime"`
	Attributes   []otelAttribute `json:"attributes"`
}

// otelAttribute is a string key/value pair carrying deterministic identifiers.
type otelAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// otelMetric is one offline counter with an integer virtual time attribute.
type otelMetric struct {
	Name        string          `json:"name"`
	Value       uint64          `json:"value"`
	VirtualTime string          `json:"virtualTime"`
	Attributes  []otelAttribute `json:"attributes"`
}

// otelLogRecord is one verdict or finding log entry with virtual time.
type otelLogRecord struct {
	SeverityText string          `json:"severityText"`
	Body         string          `json:"body"`
	VirtualTime  string          `json:"virtualTime"`
	Attributes   []otelAttribute `json:"attributes"`
}
