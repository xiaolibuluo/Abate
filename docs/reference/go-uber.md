# Go 社区规范规则（Uber Go Style Guide）— 归档参考

> 来源：https://github.com/uber-go/guide （社区，M3/M5 规则引擎落地依据）
> 归档日期：2026-11-08

## 包名
- 全小写、无大写/下划线、不复数（`net/url` 非 `net/urls`）；禁用 `common/util/shared/lib`。

## 命名
- 测试函数可含下划线（`TestMyFunc_What`）。
- 导出错误变量 `Err` 前缀，未导出 `err` 前缀；自定义错误类型 `Error` 后缀。
- 未导出全局变量/常量前缀 `_`（`_defaultPort`）。
- 禁止用 `error`/`string` 等预声明标识符命名（防 shadow）。
- 导入别名仅在包名冲突或不符路径末段时使用。

## 函数 / 结构体组织
- `import/const/var/type` 用组语法 `()` 聚类；导入分组（标准库 / 其他）。
- 函数按调用顺序、按接收者分组；`newXYZ()` 紧随类型定义。
- **禁止嵌入 Mutex**（用字段 `mu sync.Mutex`）；**公共结构体禁止嵌入类型**（用字段+委托）。
- 结构体初始化用字段名（go vet 强制）；全零值用 `var x T`；引用用 `&T{}`。
- 减少嵌套：先处理错误/特殊条件并 return。
- nil 切片合法：返回 `nil` 而非 `[]T{}`；用 `len(s)==0` 判空。

## 错误处理
- 静态串 `errors.New`；动态串 `fmt.Errorf`（`%w` 包装保留底层错误）。
- 错误仅处理一次：禁止同时日志又返回。
- 类型断言必须用 comma-ok。
- 禁止 panic（生产代码）；仅不可恢复或初始化失败可用。

## 并发
- 零值 Mutex 有效（`var mu sync.Mutex`）；不嵌入。
- 通道大小通常 1 或 0。
- 禁止无停止信号的 goroutine；`init()` 禁起 goroutine；避免可变全局（用依赖注入）。

## 测试
- 表驱动测试（`tests` 切片 + `tt` 用例 + `t.Run`）。
- `t.Parallel()` 必须显式拷贝循环变量。

## Do / Don't 速查
- **Do**：边界处复制接收的 slice/map；defer 释放；编译期校验接口（`var _ T = (*X)(nil)`）；`time.Time`/`Duration`；序列化结构体加 `json:`/`yaml:` 标签；`make` 指定容量；热路径用 `strconv`。
- **Don't**：指针到接口；公共结构体嵌入类型；内置标识符命名；生产 panic；无停止信号 goroutine；可变全局；返回 `[]int{}` 代替 nil。
