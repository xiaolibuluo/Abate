# C# 官方 / 社区规范规则（.NET Framework Design Guidelines + StyleCop）— 归档参考

> 来源：https://learn.microsoft.com/en-us/dotnet/standard/design-guidelines/ （官方）
> 社区：https://github.com/DotNetAnalyzers/StyleCopAnalyzers （StyleCop 规则 SAxxxx）
> 归档日期：2026-11-08　用途：M5 C# 前端规则引擎落地依据

## Naming Guidelines（官方）
- 程序集/命名空间/类型/成员命名一致；类型用 PascalCase；方法/属性 PascalCase；局部变量 camelCase；常量 PascalCase 或 UPPER；接口 `I` 前缀；异常 `Exception` 后缀。

## Type Design（官方）
- 静态类 vs 抽象类；接口/枚举/结构体选择；优先组合。

## Member Design（官方）
- 属性/方法/构造函数/字段/事件/运算符/参数设计；属性不应有重副作用；方法参数用最具体类型。

## Exceptions（官方）
- 抛异常而非返回错误码；自定义异常继承 Exception；不要吞异常。

## Usage Guidelines（官方）
- 数组/属性/集合/序列化/相等运算符重载的正确用法。

## 社区（StyleCop Analyzers）
- SA1xxx 文档（文件/类/方法头注释）；SA11xx 可读性；SA12xx 排序（成员/using 顺序，如 SA1202 可配置可访问性排序）；SA13xx 可维护性；SA14xx 布局；SA15xx 间距；SX1xxx 命名（如 SX1309 字段前缀 `_`）。

## 工具
- Roslyn Analyzer（StyleCop.Analyzers）按规则 ID 定制严重级；EditorConfig 配置。
