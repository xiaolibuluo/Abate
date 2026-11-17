/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// Strategy selects how an experiment is rolled back.
type Strategy string

// Rollback strategy constants.
const (
	StrategyAuto    Strategy = "auto"
	StrategyManual  Strategy = "manual"
	StrategyTimeout Strategy = "timeout"
)
