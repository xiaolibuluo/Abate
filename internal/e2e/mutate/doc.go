/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package mutate defines the four mutation classes, the operator registry, the
// in-memory edit application and the explicit equivalent-mutation set. Every
// operator works on an in-memory byte buffer only; nothing is written back to
// the read-only upstream source trees, and no file is selected by a random
// generator.
package mutate
