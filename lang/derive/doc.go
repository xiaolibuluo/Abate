/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package derive turns compiled rules into M-family mutation operators and
// runs them through the deterministic mutation kernel.
//
// The five-row derivation table is the single source of truth that maps a rule
// class to an operator class, a lexical token pair and an AIR node kind. Killing
// is a static AIR-level approximation: it never starts a subprocess.
package derive
