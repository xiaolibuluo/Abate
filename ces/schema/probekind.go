/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// ProbeKind enumerates the four probe implementations available in CES v1.
type ProbeKind string

// Probe kind constants.
const (
	ProbeHTTP   ProbeKind = "http"
	ProbeShell  ProbeKind = "shell"
	ProbeMetric ProbeKind = "metric"
	ProbeProm   ProbeKind = "prom"
)
