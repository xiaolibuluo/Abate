# ADR-YAML-001：自研严格 YAML 子集解析器

- 状态：已采纳（Accepted）
- 日期：2026-11-16
- 决策来源：高见远（架构师）`docs/refactor/DESIGN.md` §2.2
- 落盘：寇豆码（工程师）
- 关联：TODO P1-03、P1-01（ADR 部分）

> 本文记录架构师已作出的技术决策，不含手写签名；任何法律结论均不在本文范围。

---

## 1. 背景

`gopkg.in/yaml.v3 v3.0.1` 是 CES 实验文档的唯一生产解析器。合规重构要求零第三方依赖，
而 CES 文档语法面很窄，可用自研严格子集替代，无需实现完整 YAML 规范。

## 2. 实际调用面（读码事实）

| 使用形态 | 代码位置 | 性质 |
|---|---|---|
| `yaml.NewDecoder(r).KnownFields(true).Decode(&exp)` | `ces/parse/parse.go:49-51` | **唯一生产解析路径** |
| `yaml.Unmarshal` / `yaml.Marshal` | `cmd/gen-tck/{scenarios,emission}.go`、`ces/tck/{case,driver,tck_test}.go`、`ces/tck/example/builder.go`、`ces/schema/*_test.go` | 工具与测试；生成器/TCK/roundtrip |
| `encoding.TextMarshaler` / `TextUnmarshaler` | `ces/schema/duration.go`（`VirtualDuration`） | 经 text 接口编解码 `1s`/`250ms` |
| tag 形态 | `yaml:"name"` / `yaml:"name,omitempty"` | 仅这两种；无自定义 tag、无 `inline` |
| JSON 互操作 | `ces/schema/schema_test.go:165-195`：`yaml.Unmarshal(jsonBytes, &v)` | JSON 是 YAML 子集；推荐改走 `encoding/json` |

## 3. 支持语法子集（白名单）

| 类别 | 支持 | 备注 |
|---|---|---|
| 结构 | 缩进映射 `key: value`、块序列 `- item`、嵌套组合、空值 `null`/空行、行注释 `#` | 与 `ces/parse/testdata/ok-*.yaml` 一致 |
| 标量 | 双引号/单引号字符串、裸字符串、整数、浮点、布尔 `true/false` | 严格 YAML 1.2 布尔；`yes/no/on/off` 按裸字符串处理并同步 golden |
| 组合值 | `map[string]string`（`selector`/`labels`/`params`）、`[]string`（`targets`）、`[]struct` | 见 `ces/schema/*.go` tag |
| 类型接口 | `encoding.TextUnmarshaler` / `TextMarshaler` | `VirtualDuration` 依赖 |
| 严格模式 | `KnownFields(true)`：未知字段报错，错误文本含 `unknown field` | 保持 `RuleUnknownField` 分类 |
| 多文档 | 单次 `Decode`；`---` 分隔拒绝并报 `RuleMalformedDocument` | `parse.go` 只 Decode 一次 |

## 4. 明确不支持（写进包注释）

锚点 `&a` / 别名 `*a` / 合并键 `<<`、块标量 `|` / `>`、自定义 tag `!!`、复杂/非字符串 key、
流式多文档 Decode、YAML 1.1 隐式类型转换。

> 这些语法在 CES 语料中未出现；若后续需要，须走「扩展子集 + 新 golden」流程，不得静默放宽。

## 5. 实现方案

新包 `internal/yamlmini/`，公开 API：

```go
func Unmarshal(data []byte, v any) error
func Marshal(v any) ([]byte, error)
func NewDecoder(r io.Reader) *Decoder
func (d *Decoder) KnownFields(enable bool) *Decoder
func (d *Decoder) Decode(v any) error
```

内部三段式：`lexer`（按行/缩进/引号切 token）→ `node`（`map[string]*Node` / `[]*Node` / scalar 的有序树）
→ `bind`（reflect + tag + `TextUnmarshaler` 绑定，严格模式对未知 key 返回含 `unknown field` 的错误）。

**Marshal 必须确定性**：按结构体字段声明顺序输出，map 键排序，保证 roundtrip 逐字节稳定。

## 6. 验收基线（不引入任何第三方源码）

1. `ces/parse/testdata/` 全部 `ok-*` 通过、`bad-*` 全部按预期规则拒绝；
2. TCK parse 68 例 / roundtrip 50 例 / validate 46 例全绿；
3. `schema_test.go` 的 JSON↔YAML roundtrip 改走 `encoding/json` 后通过；
4. 迁移前用 yaml.v3 **仅生成 golden 输出文本**作为历史观测值，标注「历史观测值，非当前依赖」，不落地库源码。

## 7. 风险与缓解

| 风险 | 等级 | 缓解 |
|---|---|---|
| 子集漏掉 testdata/TCK 中某种语法 | 中高 | 60 ok + 90 bad + TCK 164 例强回归；不支持语法写进包注释；迁移前 golden 对拍 |
| JSON 互操作回归 | 中 | 改走标准库 `encoding/json`，不依赖自研解析器兼容 JSON |

## 8. 后果

- 正向：编译依赖清零；解析器行为完全可审计、逐字节确定。
- 反向：维护一个项目自有解析器；扩展语法必须显式走流程，不能静默放宽。
