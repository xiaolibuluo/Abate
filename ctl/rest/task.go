/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

// Task is one asynchronous check or experiment recorded by the control plane.
//
// It carries no wall-clock field: status transitions and counters are the only
// observable progress, keeping task records deterministic and offline.
type Task struct {
	ID             string      `json:"id"`
	Kind           string      `json:"kind"`
	Language       string      `json:"language"`
	Format         string      `json:"format"`
	Status         TaskStatus  `json:"status"`
	FingerprintHex string      `json:"fingerprint,omitempty"`
	Summary        summaryView `json:"summary,omitempty"`
	Detail         string      `json:"detail,omitempty"`
	Report         []byte      `json:"-"`
	Cancelled      bool        `json:"cancelled"`
}

// summaryView is the JSON-facing projection of a check summary.
type summaryView struct {
	Files      int            `json:"files"`
	Rules      int            `json:"rules"`
	Total      int            `json:"total"`
	BySeverity map[string]int `json:"bySeverity,omitempty"`
}
