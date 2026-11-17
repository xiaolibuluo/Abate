/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package bench holds the XCU-PRF performance budget table and the deterministic
// benchmark suite that reports real measured numbers without hard assertions.
//
// 参考 08 分册 XCU-PRF 预算表 P1-P15，改进点=预算阈值以整数纳秒或每秒计数
// 表达，禁止浮点时间与硬断言（性能抖动不判 FAIL），只记录真实数值。
package bench

import "strconv"

// Budget is one quantified performance budget item of the XCU-PRF table.
type Budget struct {
	// ID is the stable budget identifier, P1 through P15.
	ID string

	// Metric names the measured quantity.
	Metric string

	// Threshold is the integer bound in the unit declared by Unit.
	Threshold int64

	// Unit is the measurement unit, one of events/s, draws/s, ns or bytes/s.
	Unit string

	// Owner is the module accountable for the budget.
	Owner string
}

// Budgets returns the fifteen XCU-PRF budget rows in fixed identifier order.
func Budgets() []Budget {
	return []Budget{
		{"P1", "scheduler event throughput", 100000, "events/s", "KRN-SCH"},
		{"P2", "PRNG draw throughput", 100000000, "draws/s", "SIM-RNG"},
		{"P3", "virtual network propagation of 100k packets P99", 1000000000, "ns", "SIM-NET"},
		{"P4", "virtual scheduling of 10k tasks throughput", 100000, "events/s", "SIM-SCHED"},
		{"P5", "archive Load+Verify of 1M frames P99", 10000000000, "ns", "SIM-REC"},
		{"P6", "AIR serialization of 1M nodes P99", 500000000, "ns", "LNG-AIR"},
		{"P7", "AIR deserialization of 1M nodes P99", 700000000, "ns", "LNG-AIR"},
		{"P8", "frontend parse throughput", 1000000, "bytes/s", "LNG-*"},
		{"P9", "rule matching of 100 rules P99", 50000000, "ns", "LNG-RULE"},
		{"P10", "HTML report of 100k findings P99", 5000000000, "ns", "RPT-HTML"},
		{"P11", "plugin invocation overhead P99", 50000, "ns", "WAS-HOST"},
		{"P12", "plugin signature verification P99", 10000000, "ns", "WAS-SIG"},
		{"P13", "CES parse of a 100KB document P99", 10000000, "ns", "CES-PARSE"},
		{"P14", "CES validation of a 100KB document P99", 5000000, "ns", "CES-VALID"},
		{"P15", "CI fast-lane total duration P95", 600000000000, "ns", "XCU-CI"},
	}
}

// BudgetByID resolves one budget row, reporting whether it exists.
func BudgetByID(id string) (Budget, bool) {
	for _, budget := range Budgets() {
		if budget.ID == id {
			return budget, true
		}
	}
	return Budget{}, false
}

// String renders one budget row in a fixed, deterministic field order.
func (b Budget) String() string {
	return b.ID + "|" + b.Owner + "|" + b.Metric + "|" +
		strconv.FormatInt(b.Threshold, 10) + b.Unit
}
