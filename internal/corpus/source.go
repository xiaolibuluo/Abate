/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

// Source labels the read-only track an entry belongs to. It is report metadata
// only and never enters a control-flow condition.
type Source string

// SourceInTree marks the pre-existing in-tree upstream source tree.
const SourceInTree Source = "in_tree"

// SourceValidation marks the independently downloaded validation corpus.
const SourceValidation Source = "validation"

// SourceSelfResearched marks first-party corpus bytes written from scratch by
// this project. It is report metadata only and never enters a control-flow
// condition.
const SourceSelfResearched Source = "self_researched"
