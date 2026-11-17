/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"strings"

	"abate/lang/frontkit"
)

// jsTokenSpelling 拼接一段 token 的真实源码拼写，供非 lexical 容器节点绑定.
//
// 正则字面量等容器节点不会被 cover 投影成逐字符叶节点；只有把整段真实
// spelling 写入节点 attribute，其内部单字符变异才会改变 AIR-ID。
func jsTokenSpelling(toks []frontkit.Token, lo, hi int) string {
	if lo < 0 {
		lo = 0
	}
	if hi > len(toks) {
		hi = len(toks)
	}
	var out strings.Builder
	for i := lo; i < hi; i++ {
		out.WriteString(toks[i].Text)
	}
	return out.String()
}
