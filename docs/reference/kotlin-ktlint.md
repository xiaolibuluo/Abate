# Kotlin 社区规范规则（ktlint 标准规则集）— 归档参考

> 来源：ktlint（pinterest.github.io/ktlint，主站无可读内容，规则基线以下列标准集为准）
> 归档日期：2026-11-08　用途：M5 Kotlin 前端规则引擎落地依据

## ktlint-official 标准规则（基线）
- 命名：`class-naming`（UpperCamelCase）、`function-naming`（camelCase）、`property-naming`（camelCase）、`package-naming`（小写无下划线）、`enum-entry-naming`。
- 格式：`indentation`（4 空格，禁 tab）、`curly-spaces`（大括号空格）、`no-empty-first-line-in-class-body`、`wrapping`（链式调用换行）、`blank-line-before-declaration`、`blank-line-between-when-branches`。
- 空格：`no-whitespace-in-continuation-indent`、运算符/分隔符周围空格（`spacing-around-keyword`、`spacing-around-operators`、`spacing-around-comma`）。
- 导入：`import-ordering`（字典序分组）、`no-unused-imports`、`no-wildcard-imports`。
- 其它：`no-empty-first-line-in-class-body`、`modifier-order`、`annotation`、`string-template`、`range-spacing`。

## ktlint-experimental（谨慎启用）
- `function-signature`、`class-signature`、`property-naming` 等更严格签名格式。

## 说明
- 主站不可达，规则项取自 ktlint 公开标准规则集文档；M5 落地时以 ktlint 官方规则 ID 映射为 AIR 模式匹配。
