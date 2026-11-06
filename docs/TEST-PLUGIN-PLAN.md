<!--
Abate —— 测试插件生态规划
Copyright (C) 2026 小黎的部落（www.qingdujian.com）
AGPLv3; 商业用途须事先获得小黎的部落书面授权
-->
# 测试插件生态规划 | Test Plugin Ecosystem Plan

> 23 种测试插件，按需安装，矩阵编排，CI 强制。

## 1. 设计原则

- 一类一文件：每种测试插件独立成文件，单文件 <=300 行
- 按需安装：插件以 WASM 模块形式分发，abate plugin install <name> 安装
- 统一接口：所有测试插件实现 TestPlugin 接口
- 矩阵编排：abate test --matrix 串行运行全部插件，输出矩阵报告
- 100% 覆盖：每个插件自身也需要测试覆盖

## 2. TestPlugin 接口设计

```go
type TestPlugin interface {
    Capability() core.Capability
    Setup(ctx context.Context) (cleanup func(),  err error)
    Run(ctx context.Context, target string) TestResult
    Teardown(ctx context.Context) error
}

type TestResult struct {
    Name     string
    Passed   bool
    Duration time.Duration
    Details  string
    Metrics  map[string]float64
}
```

## 3. 23 种测试插件清单

### 第一组：功能测试（5 个）

| # | 插件名 | 文件 | 描述 |
|---|--------|------|------|
| 1 | unit-test | test/unit_test.go | 函数级隔离测试 |
| 2 | property-test | test/property_test.go | 随机输入验证不变式 |
| 3 | integration-test | test/integration_test.go | 多包协作场景验证 |
| 4 | concurrency-test | test/concurrency_test.go | -race 检测数据竞争 |
| 5 | compatibility-test | test/compatibility_test.go | Go 1.21+ 兼容性 |

### 第二组：性能与负载测试（4 个）

| # | 插件名 | 文件 | 描述 |
|---|--------|------|------|
| 6 | perf-baseline | test/perf_baseline_test.go | benchmark 对比基线 |
| 7 | backpressure-test | test/backpressure_test.go | 限流下行为验证 |
| 8 | stress-test | test/stress_test.go | 超负载稳定性 |
| 9 | soak-test | test/soak_test.go | 长时间内存泄漏检测 |

### 第三组：健壮性与安全测试（8 个）

| # | 插件名 | 文件 | 描述 |
|---|--------|------|------|
| 10 | fuzz-test | test/fuzz_test.go | 随机/畸形输入不崩溃 |
| 11 | conflict-test | test/conflict_test.go | 并发注入互斥故障不死锁 |
| 12 | defect-test | test/defect_test.go | 注入已知缺陷验证检测率 |
| 13 | crash-test | test/crash panic 后恢复验证 |
| 14 | vuln-test | test/vuln_test.go | 已知 CVE 模式扫描 |
| 15 | security-test | test/security_test.go | 输入注入/路径穿越检测 |
| 16 | adversarial-test | test/adversarial_test.go | 恶意参数对抗验证 |
| 17 | fault-tolerance-test | test/fault_tolerance_test.go | 部分依赖失败后降级验证 |

### 第四组：工程检查（6 个）

| # | 插件名 | 文件 | 描述 |
|---|--------|------|------|
| 18 | lint-intensive | test/lint_intensive_test.go | golangci-lint 全规则 |
| 19 | static-analysis | test/static_analysis_test.go | go vet + gosec |
| 20 | dynamic-analysis | test/dynamic_analysis_test.go | -race + -cover + 内存检测 |
| 21 | cicd-check | test/cicd_check_test.go | 模拟完整流水线 |
| 22 | dual-build-verify | test/dual_build_verify_test.go | debug+release 编译 |
| 23 | dual-mode-run | test/dual_mode_run_test.go | debug+release 运行测试 |

## 4. 实施顺序（28 步串行）

1. test/plugin.go - TestPlugin 接口 + TestResult
2. test/registry.go - 测试插件注册表
3. test/runner.go - 矩阵编排器
4-8. 第一组功能测试插件（5 个）
9-12. 第二组性能负载插件（4 个）
13-20. 第三组健壮安全插件（8 个）
21-26. 第四组工程检查插件（6 个）
27. 更新 .github/workflows/redline.yml
28. 提交推送 GitHub

---

Copyright (C) 2026 小黎的部落（www.qingdujian.com）