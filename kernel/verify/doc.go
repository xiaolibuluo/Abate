/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package verify turns an observation stream into experiment verdicts.
//
// Every decision is a pure function of the observations and the steady-state
// hypotheses: replaying a recorded stream always yields byte-identical
// verdicts. No I/O and no real clock are involved.
package verify
