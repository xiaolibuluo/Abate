# 自研变异语义冻结 golden

本目录存放 `mutation-semantics.golden`，由本仓 `kernel/mutation/pitparity` 引擎对同构夹具的
真实判定生成，**不是上游官方产物**，也不包含任何上游源码。

- 生成器：`kernel/mutation/pitparity/gen_golden.go`（`//go:build ignore`）
- 复跑命令：`go run kernel/mutation/pitparity/gen_golden.go`
- 生成时点：T01（清除动作发生前）
- 用途：`scripts/pit-parity.sh` 已改为「自研语义 golden 回归」口径，阈值
  `MIN_RATE=100.00%` 保持不变；竞品对拍分数为 `NOT_MEASURED`，不再宣称「官方对拍 100%」。

## 口径声明

1. golden 中的判定值来自本仓变异核（registry + engine + verifier）与语义见证，非外部引擎实跑；
2. golden 为历史观测快照，后续变更须通过生成器重跑并记录原因，不得手改数值；
3. ⚖️ 上游产物的历史并入义务需法务确认（TODO L04），本文不作法律结论。
