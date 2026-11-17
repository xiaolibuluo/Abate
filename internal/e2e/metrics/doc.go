/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package metrics defines the four-way mutation columns, precision rows,
// robustness counters, drilldown records and the pure aggregation pass. Every
// reported number is derived from drilldown records so a runner can recompute
// the same values with zero difference. The four columns stay separate types;
// this package deliberately provides no weighted, offsetting or total score.
package metrics
