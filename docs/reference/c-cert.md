# C/C++ 官方规范规则（SEI CERT C/C++ Secure Coding）— 归档参考

> 来源：https://wiki.sei.cmu.edu/confluence/display/c/SEI+CERT+C+Coding+Standard （官方安全编码基线）
> 归档日期：2026-11-08　用途：M5 C/C++ 前端规则引擎安全规则来源

## 规则分组（C）
- PRE 预处理 / DCL 声明与初始化 / EXP 表达式 / INT 整数 / FLP 浮点 / ARR 数组
- STR 字符与字符串 / MEM 内存管理 / FIO 输入输出 / ENV 环境 / SIG 信号
- ERR 错误处理 / API 应用编程接口 / CON 并发 / MSC 杂项 / POS POSIX / WIN Windows

## 核心安全主题（规则集设计依据）
- 内存安全：禁止越界访问、悬垂指针、释放后使用（USE AFTER FREE）。
- 整数安全：符号/宽度转换、整数溢出（INT30/INT32）。
- 字符串：禁止无界拷贝（`strcpy`→`strncpy`/边界检查）、空终止保障。
- 内存管理：配对分配/释放、禁止双重释放、资源 RAII 式管理。
- 输入验证：不可信输入边界检查前置（FIO/ENV）。
- 并发：数据竞争、原子性、禁止在信号处理中调用非异步安全函数（SIG31/SIG34）。
- 错误处理：禁止忽略返回值、禁止 `setjmp/longjmp` 跨越初始化。

## 说明
- CERT 区分 Rules（规范性）与 Recommendations（指导性）。
- 提供 Juliet 测试套件（61,387 用例）供静态分析工具覆盖度验证。
