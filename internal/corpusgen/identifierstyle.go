/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// IdentifierStyle selects the spelling convention of a long code identifier.
type IdentifierStyle string

// StyleCamel selects camel-case or Pascal-case identifiers.
const StyleCamel IdentifierStyle = "camel"

// StyleSnake selects snake-case identifiers.
const StyleSnake IdentifierStyle = "snake"
