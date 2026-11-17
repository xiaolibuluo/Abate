/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package rule defines the language-neutral rule model and the AIR structural
// pattern matching engine shared by the Abate checker and the mutation pipeline.
//
// Rules are declared as table-driven Specifications. A Registry validates and
// compiles each Specification into a Rule whose Check method accepts only an
// AIR unit, so no language-specific syntax type can enter the public contract.
package rule
