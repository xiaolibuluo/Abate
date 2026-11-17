# Go 官方规范规则（Effective Go）— 归档参考

> 来源：https://go.dev/doc/effective_go （官方，M3/M5 规则引擎落地依据）
> 归档日期：2026-11-08　用途：自研 Go 规则集官方基线

## 命名
- 包名：全小写单字，无下划线、无 mixedCaps（`lower case, single-word`）。
- 可见性：首字母大写导出，小写私有。
- Getter 不加 `Get` 前缀（`Owner` 而非 `GetOwner`）；setter 用 `SetOwner`。
- 单方法接口以 `-er` 后缀（`Reader`/`Writer`/`String`）。
- 多词标识符用 `MixedCaps`，**禁止下划线**。

## 格式化
- 必须 gofmt（tab 缩进）。Go 无行长度硬限；行长用换行+额外 tab。
- 控制结构 `if/for/switch/select` 开括号 `{` **不得**另起一行（否则词法器插分号）。

## 注释 / 文档
- 顶层声明前紧邻注释即文档注释（godoc 来源）。

## 控制结构
- `if`/`switch` 支持初始化语句；成功路径以 `return` 结束可省略 `else`。
- 仅 `for` 循环；`range` 遍历切片/映射/通道/字符串（UTF-8）。
- `switch` 无自动 fallthrough，需显式列 case。

## 函数
- 多返回值（`(n int, err error)`）；可用命名返回参数。
- `defer` 管理资源（文件/锁），LIFO，参数在 defer 执行时求值。
- `init()` 仅做状态设置/校验。

## 内存分配
- `new(T)` 返回 `*T` 零值；`make` 仅用于 slice/map/channel，返回值非指针。
- 数组是值类型；优先切片。Map 用 comma-ok 区分缺失与零值。

## 接口 / 类型
- 隐式实现；接口嵌入组合；结构体匿名嵌入获方法。
- 指针 vs 值接收者：值接收者可被指针/值调用，指针接收者只能被指针调用。
- 构造函数若类型仅为实现接口，返回接口值而非具体类型。

## 并发
- 「不要通过共享内存通信，而通过通信共享内存」。
- goroutine + channel；`select` 多路；并发≠并行。

## Do / Don't
- **Do**：gofmt 全部；包名小写单字；早 return 省 else；defer 管资源；comma-ok；零值可用。
- **Don't**：getter 不加 Get；命名不用下划线；开括号不换行；switch 不依赖 fallthrough；**绝不忽略 error**；库函数避免 panic；类型断言不用崩溃式；不共享内存通信。
