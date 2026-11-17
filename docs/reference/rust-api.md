# Rust 官方规范规则（Rust API Guidelines, C-* 检查清单）— 归档参考

> 来源：https://rust-lang.github.io/api-guidelines/checklist.html （官方，M5 规则引擎落地依据）
> 归档日期：2026-11-08

## Naming
- C-CASE RFC 430 大小写；C-CONV `as_/to_/into_` 约定；C-GETTER getter 命名；C-ITER `iter/iter_mut/into_iter`；C-ITER-TY 迭代器类型名匹配；C-FEATURE 特性名无占位词；C-WORD-ORDER 词序一致。

## Interoperability
- C-COMMON-TRAITS  eagerly 实现 Copy/Clone/Eq/PartialEq/Ord/PartialOrd/Hash/Debug/Display/Default；C-CONV-TRAITS 用 From/AsRef/AsMut；C-COLLECT 集合实现 FromIterator/Extend；C-SERDE 实现 Serialize/Deserialize；C-SEND-SYNC 尽量 Send/Sync；C-GOOD-ERR 错误类型良好；C-NUM-FMT 提供 Hex/Octal/Binary；C-RW-VALUE 读写函数按值取 `R: Read`/`W: Write`。

## Macros
- C-EVOCATIVE 输入语法呼应输出；C-MACRO-ATTR 与属性组合；C-ANYWHERE 项宏随处可用；C-MACRO-VIS 支持可见性；C-MACRO-TY 类型片段灵活。

## Documentation
- C-CRATE-DOC 充分且含示例；C-EXAMPLE 所有项有 rustdoc 示例；C-QUESTION-MARK 用 `?` 非 unwrap；C-FAILURE 文档含 error/panic/safety；C-LINK 超链接；C-METADATA Cargo.toml 全元数据；C-RELNOTES 发布说明。

## Predictability
- C-SMART-PTR 智能指针不加固有方法；C-CONV-SPECIFIC 转换置于最具体类型；C-METHOD 明确接收者用方法；C-NO-OUT 不取输出参数；C-OVERLOAD 运算符重载不意外；C-DEREF 仅智能指针实现 Deref；C-CTOR 构造函数为静态固有方法。

## Flexibility
- C-INTERMEDIATE 暴露中间结果；C-CALLER-CONTROL 调用方决定拷贝/放置；C-GENERIC 用泛型最小化参数假设；C-OBJECT 可能有用的 trait 对象安全。

## Type safety
- C-NEWTYPE 新类型提供静态区分；C-CUSTOM-TYPE 参数用类型而非 `bool`/`Option` 表义；C-BITFLAG 标志集用 bitflags 非 enum；C-BUILDER 复杂值用 builder。

## Dependability
- C-VALIDATE 函数校验参数；C-DTOR-FAIL 析构不失败；C-DTOR-BLOCK 可能阻塞析构有替代。

## Debuggability
- C-DEBUG 所有公开类型实现 Debug；C-DEBUG-NONEMPTY Debug 非空。

## Future proofing
- C-SEALED 密封 trait 防下游实现；C-STRUCT-PRIVATE 结构体私有字段；C-NEWTYPE-HIDE 新类型封装实现；C-STRUCT-BOUNDS 不重复派生 trait bound。

## Necessities
- C-STABLE 稳定 crate 的公开依赖稳定；C-PERMISSIVE crate 及依赖为宽松许可证。
