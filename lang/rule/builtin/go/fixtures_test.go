/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "strings"

// fixture holds one positive and one negative source fragment for a rule.
type fixture struct {
	Pos string
	Neg string
}

// allFixtures joins the three specification-family fixture tables.
func allFixtures() map[string]fixture {
	out := make(map[string]fixture, 128)
	for _, table := range []map[string]fixture{
		egFixtures(), crcFixtures(), vetFixtures(),
	} {
		for id, item := range table {
			out[id] = item
		}
	}
	return out
}

// fixtureSource wraps a fragment into a parsable Go source file.
func fixtureSource(fragment string) []byte {
	if strings.Contains(fragment, "package ") {
		return []byte(fragment)
	}
	return []byte("package p\n" + fragment + "\n")
}
