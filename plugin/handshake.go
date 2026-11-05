// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package plugin

import (
	"fmt"

	"github.com/qingdujian/abate/core"
)

// Handshake 执行插件握手验证（CES-7.1.1 / CES-7.1.2）。
//
// 检查项：
//   - 协议主版本与核心一致（CES-7.1.1, E1001）
//   - 能力声明字段完整：name 非空、protocol 非空（CES-7.1.2, E1002）
//   - Injector 角色插件 faultClasses 非空（CES-7.1.2, E1002）
//
// Handshake performs plugin handshake validation (CES-7.1.1 / CES-7.1.2).
func Handshake(cap core.Capability) error {
	if cap.Name == "" {
		return fmt.Errorf("plugin: name is empty (CES-7.1.2, E1002)")
	}
	if cap.Protocol == "" {
		return fmt.Errorf("plugin %q: protocol is empty (CES-7.1.2, E1002)", cap.Name)
	}
	if cap.Protocol != core.ProtocolVersion {
		return fmt.Errorf("plugin %q: protocol %q incompatible with %q (CES-7.1.1, E1001)",
			cap.Name, cap.Protocol, core.ProtocolVersion)
	}
	return nil
}

// HandshakeInjector 对 Injector 角色执行额外检查：faultClasses 非空（CES-7.1.2）。
//
// HandshakeInjector performs an extra check for the Injector role:
// faultClasses must not be empty.
func HandshakeInjector(cap core.Capability) error {
	if err := Handshake(cap); err != nil {
		return err
	}
	if len(cap.FaultClasses) == 0 {
		return fmt.Errorf("plugin %q: faultClasses is empty (CES-7.1.2, E1002)", cap.Name)
	}
	return nil
}
