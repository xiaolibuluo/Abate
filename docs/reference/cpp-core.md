# C++ 社区规范规则（C++ Core Guidelines, isocpp）— 归档参考

> 来源：https://isocpp.github.io/CppCoreGuidelines/CppCoreGuidelines （社区，M5 规则引擎落地依据）
> 归档日期：2026-11-08

## P 哲学
- P.1 直接在代码表达想法；P.3 明确意图（范围 for 优于裸索引）；P.4 理想静态类型安全；P.5 优先编译期检查；P.8 不泄漏资源（RAII）；P.10 优先不可变；P.11 封装杂乱构造（vector 代替裸 malloc）。

## I 接口
- I.1 接口明确（避免隐式全局控制）；I.2 避免非常量全局；I.3 避免单例；I.4 接口精确强类型（避免 `void*`）；I.5/I.6 声明前置条件（GSL `Expects()`）；I.7 后置条件（`Ensures()`）；I.11 不通过裸指针/引用转移所有权；I.12 必非空声明 `not_null`；I.13 不用单指针传数组（用 `span`）；I.23 参数少于 4 个。

## F 函数
- F.1 有意义操作打包为函数；F.3 简短简单；F.4 可能编译期求值用 `constexpr`；F.6 不抛用 `noexcept`；F.7 通用用 `T*`/`T&` 而非智能指针；F.15-19 参数传递（in 值/const 引用，in-out 非 const 引用，move 用 `X&&`，forward 用 `TP&&`）；F.20-21 多返回值用 struct/返回；F.26 `unique_ptr` 转移所有权；F.27 `shared_ptr` 共享；F.43 不返回局部指针/引用；F.49 不返回 `const T`。

## 其余组（目录存在，需后续补全抓取）
- C 类层次 / EN 枚举 / R 资源管理 / ES 表达式语句 / Per 性能 / CP 并发 / E 错误处理 / Con 常量 / T 模板 / CPL C 风格 / SF 源文件。

## 关键 RAII / 所有权（核心）
- 资源所有权用智能指针；数组用 `span<T>`；C 风格字符串用 `zstring`/`not_null<zstring>`。
