// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package web

import (
	"net/http"

	"github.com/qingdujian/abate/hub"
)

// hubSubcodesHandler 暴露 ChaosHub 子编码注册表（CES-B.1）。
// hubSubcodesHandler exposes the ChaosHub sub-code registry (CES-B.1).
func hubSubcodesHandler(w http.ResponseWriter, _ *http.Request) {
	reg := hub.NewSeedRegistry()
	writeJSON(w, http.StatusOK, map[string]any{
		"version": reg.Version(),
		"count":   reg.Count(),
		"codes":   reg.AllCodes(),
	})
}
