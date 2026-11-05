// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package injector

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/qingdujian/abate/core"
)

// subCodeParams 是附录 B 子编码的默认参数解析表。
var subCodeParams = map[string]map[string]string{
	"F1.1": {"delay": "100ms"},
	"F1.2": {"delay": "100ms", "jitter": "10ms"},
	"F1.3": {"timeoutMs": "5000"},
	"F2.1": {"signal": "SIGTERM"},
	"F2.2": {"signal": "SIGTERM", "gracePeriodMs": "5000"},
	"F3.1": {"loss": "50", "direction": "to"},
	"F3.2": {"domain": "example.com", "mode": "random"},
	"F3.3": {"peer": "0.0.0.0/0", "direction": "to"},
	"F3.4": {"cidr": "10.0.0.0/8"},
	"F4.1": {"utilization": "80"},
	"F4.2": {"utilization": "80", "fillRate": "10"},
	"F4.3": {"utilization": "90"},
	"F4.4": {"count": "1000"},
	"F4.5": {"count": "100"},
	"F4.6": {"iops": "1000"},
	"F5.1": {"bitRate": "1", "pattern": "random"},
	"F5.2": {"percent": "10"},
	"F5.3": {"percent": "10"},
	"F6.1": {"offsetMs": "5000"},
	"F6.2": {"offsetMs": "-5000"},
}

// Router 是注入器多路复用器，按故障类路由到对应注入器。
// Router is an injector multiplexer that routes by fault class.
type Router struct {
	mu       sync.Mutex
	injectors map[core.FaultClass]core.Injector
	active   map[string]core.FaultClass
}

// NewRouter 构造空路由器。
func NewRouter() *Router {
	return &Router{
		injectors: make(map[core.FaultClass]core.Injector),
		active:   make(map[string]core.FaultClass),
	}
}

// Register 注册注入器，按其 Capability 中的 FaultClasses 建立路由。
func (r *Router) Register(inj core.Injector) error {
	cap := inj.Capability()
	if len(cap.FaultClasses) == 0 {
		return fmt.Errorf("router: injector %s has no fault classes", cap.Name)
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, fc := range cap.FaultClasses {
		if _, exists := r.injectors[fc]; exists {
			return fmt.Errorf("router: fault class %s already registered", fc)
		}
		r.injectors[fc] = inj
	}
	return nil
}

// Capability 返回聚合能力声明。
func (r *Router) Capability() core.Capability {
	classes := make([]core.FaultClass, 0, len(r.injectors))
	for fc := range r.injectors {
		classes = append(classes, fc)
	}
	return core.Capability{
		Name:         "router",
		Version:      "1.0.0",
		Protocol:     core.ProtocolVersion,
		FaultClasses: classes,
		TargetKinds:  []core.TargetKind{core.TargetService, core.TargetInstance, core.TargetDependency},
	}
}

// Inject 按故障类路由到对应注入器（CES-2.4.1 幂等）。
func (r *Router) Inject(ctx context.Context, e core.Experiment) (core.Injection, error) {
	fc := e.Fault.Class
	r.mu.Lock()
	inj, ok := r.injectors[fc]
	if ok {
		r.active[e.ID] = fc
	}
	r.mu.Unlock()
	if !ok {
		return core.Injection{}, fmt.Errorf("router: no injector for fault class %s", fc)
	}
	return inj.Inject(ctx, e)
}

// Recover 路由到对应注入器恢复（CES-3.1.2 可重复调用）。
func (r *Router) Recover(ctx context.Context, in core.Injection) error {
	r.mu.Lock()
	fc, ok := r.active[in.ExperimentID]
	r.mu.Unlock()
	if !ok {
		return nil
	}
	r.mu.Lock()
	delete(r.active, in.ExperimentID)
	r.mu.Unlock()
	inj := r.injectors[fc]
	return inj.Recover(ctx, in)
}

// IsActive 报告是否有活跃注入。
func (r *Router) IsActive() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.active) > 0
}

// sortedSubCodes 是 subCodeParams 键的字典序排列，保证确定性查找。
var sortedSubCodes = func() []string {
	keys := make([]string, 0, len(subCodeParams))
	for k := range subCodeParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}()

// parseSubCode 从实验中提取子编码并填充默认参数。
func parseSubCode(e *core.Experiment) {
	subCode := string(e.Fault.Class)
	if e.Fault.Params == nil {
		e.Fault.Params = make(map[string]string)
	}
	// 精确匹配优先
	if _, ok := subCodeParams[subCode]; !ok {
		// 前缀匹配：按字典序取第一个匹配项
		for _, sc := range sortedSubCodes {
			if strings.HasPrefix(sc, subCode) && len(sc) > len(subCode) {
				subCode = sc
				break
			}
		}
	}
	if defaults, ok := subCodeParams[subCode]; ok {
		for k, v := range defaults {
			if _, exists := e.Fault.Params[k]; !exists {
				e.Fault.Params[k] = v
			}
		}
	}
}
