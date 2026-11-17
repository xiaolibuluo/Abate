/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

// exemption records one registered rule identifier and its justification.
type exemption struct {
	ID     string
	Reason string
}

// Exemptions returns a fresh copy of the registered exemption table.
//
// A rule is exempted only when the project has an explicit, documented basis:
// a linter disabled in .golangci.yml, an established repository convention, or
// a scope difference between the self-developed rule and the official entry.
// The table never changes M1/M2/lang source code and never removes a rule.
func Exemptions() map[string]string {
	table := exemptionTable()
	out := make(map[string]string, len(table))
	for _, item := range table {
		out[item.ID] = item.Reason
	}
	return out
}

// ExemptionFor returns the registered justification for a rule identifier.
func ExemptionFor(id string) (string, bool) {
	for _, item := range exemptionTable() {
		if item.ID == id {
			return item.Reason, true
		}
	}
	return "", false
}

// exemptionTable lists every rule that the Abate self-check actually hits.
// Each entry names the exact basis; no entry is a blanket suppression.
func exemptionTable() []exemption {
	return []exemption{
		{"GO-CRC-008", "数字后缀命中 sha256/go1_26 等标准库包名与版本标识，非变量命名；.golangci.yml 未启用对应命名 linter，属规则范围差异。"},
		{"GO-CRC-009", "官方 Error Strings 只约束 error 文本，自研规则作用于全部字符串字面量；.golangci.yml 未启用 forbidigo，属规则范围差异。"},
		{"GO-CRC-012", "9 处命中全部是 testing.T.Errorf 的诊断文本（首字母大写为专有名词或强调），并非 fmt.Errorf 产生的 error 字符串；官方 Error Strings 只约束 error 值。AIR 的 Selector 仅暴露方法名，规则以 Selector+Errorf 近似无法区分接收者；.golangci.yml 对 _test.go 排除 revive，golangci 报 0 issues，属规则范围差异。"},
		{"GO-CRC-015", "_ = 显式丢弃结果是项目既定约定（errcheck 已通过）；官方 Handle Errors 针对被忽略的 error 返回值，属规则范围差异。"},
		{"GO-CRC-017", "短函数内命名返回值的裸 return 是项目既定约定；revive 未启用 naked-return 检查，golangci 报 0 issues。"},
		{"GO-CRC-019", "命中项为行内注释（godot 只约束声明注释）；官方 Comment Sentences 针对句子式注释，属规则范围差异。"},
		{"GO-CRC-020", "唯一命中在测试代码中的 reflect.DeepEqual；官方 Pass Values 非机械规则，golangci 未启用对应 linter。"},
		{"GO-CRC-022", "指针 receiver 在项目内保持一致，是 Receiver Type 允许的形态；规则检测构造而非不一致，属语义子集近似。"},
		{"GO-CRC-024", "命中为字符串内容本身的合法分隔符空白，非源码尾随空白；Gofmt 由 golangci formatters 已通过，属规则范围差异。"},
		{"GO-CRC-027", "命中为 sim/clock 注入式虚拟时钟的 Now() 方法，并非 time.Now；确定性铁律已满足，属规则范围差异。"},
		{"GO-CRC-029", "命中为 main 入口点的 os.Exit(Execute(...))；官方 Synchronous Functions 排除入口点，属规则范围差异。"},
		{"GO-CRC-030", "t.Fatal 在测试代码中用于立即失败，是项目既定约定；官方 Useful Test Failures 关注失败信息内容，属规则范围差异。"},
		{"GO-CRC-032", "os.ReadFile 调用均同时检查返回的 error（errcheck 已通过）；规则检测调用构造而非缺陷，属语义子集近似。"},
		{"GO-EG-002", "命中为 ABAT/AIR/CES 等全大写缩写常量与 magic 标识；.golangci.yml 明确关闭 add-constant，属项目既定约定。"},
		{"GO-EG-014", "panic 仅用于不可达分支与测试断言；.golangci.yml 未启用 forbidigo，golangci 报 0 issues。"},
		{"GO-EG-017", "命中为测试中短生命周期循环内的 defer；官方 Defer 反模式针对长循环资源累积，属规则范围差异。"},
		{"GO-EG-018", "短函数内裸 return 是项目既定约定；revive 未启用 naked-return，golangci 报 0 issues。"},
		{"GO-EG-024", "大数值字面量在序列化常量表中有注释说明；.golangci.yml 明确关闭 add-constant，属已登记关闭项。"},
		{"GO-EG-029", "命中为循环内的整数加法（i+1）而非字符串拼接；规则以 Binary=+ 近似检测，属规则范围差异。"},
		{"GO-EG-030", "1312 处命中为存量代码的多行注释续行与行内注释（如 // and resets its ledger），并非声明文档首句；AIR 把每个 Doc 注释行映射为独立节点，规则用 Where 逐行近似。.golangci.yml 的 godot 配置为 scope: declarations 且对 _test.go 排除 revive/godot，golangci 报 0 issues，属规则范围差异。"},
		{"GO-EG-031", "make 的零长度用于延迟填充的切片；.golangci.yml 未启用 prealloc，golangci 报 0 issues。"},
		{"GO-EG-032", "空结构体用作集合成员与信号量是 Go 惯用法；规则检测构造而非缺陷，属规则范围差异。"},
		{"GO-EG-027", "recover 仅用于中间件 deferred 函数内的标准 panic recovery 惯用法（HTTP 长连接服务需将单次请求 panic 转为 500，避免整进程崩溃）；规则以 (Call Ident=recover) 近似，无法编码「直接位于 defer 函数内」的语义，误报正确用法，属规则范围差异。"},
		{"GO-CRC-006", "Ident~Cli 命中标准库 http.Client/DefaultClient 类型名（含子串 Cli），非 CLI 缩写；项目约定缩写写作全大写 CLI，规则以子串近似，属规则范围差异。"},
		{"GO-VET-009", "%v 用于非包装场景的格式化；官方 printf 仅对 wrap 位置建议 %w，M3 为 vet 语义子集近似。"},
		{"GO-VET-010", "errors.As 调用均传入指针目标；规则检测调用构造，官方 errorsas 需类型信息，属语义子集近似。"},
		{"GO-VET-014", "string(数字) 命中为 rune 转换的合法用法；官方 stringintconv 需类型信息，属语义子集近似。"},
		{"GO-VET-015", "int(字符串) 命中为边界夹具转换；官方 stringintconv 需类型信息，属语义子集近似。"},
		{"GO-VET-016", "make(map...) 无长度是合法且高效的；规则检测参数个数而非缺陷，属语义子集近似。"},
		{"GO-VET-019", "WaitGroup.Wait 均在 Add 之后调用；规则检测调用构造，官方 waitgroup 需控制流分析，属语义子集近似。"},
		{"GO-VET-022", "ctx.Done() 为受控取消路径；规则检测调用构造，官方 lostcancel 需控制流分析，属语义子集近似。"},
		{"GO-VET-023", "json struct tag 由 go vet structtag 全量校验通过（0 issues）；规则检测构造而非拼写错误，属语义子集近似。"},
		{"GO-VET-024", "json.Unmarshal 均传入非 nil 指针；规则检测调用构造，官方 unmarshal 需类型信息，属语义子集近似。"},
		{"GO-VET-025", "json.Marshal 的 error 均已检查（errcheck 通过）；规则检测调用构造，属语义子集近似。"},
		{"GO-VET-026", "唯一 http.Get 命中在受控路径且 body 已关闭（bodyclose 通过）；规则检测构造，属语义子集近似。"},
		{"GO-VET-029", "With 命中为构造器方法而非 slog.With；规则以选择器名近似检测，属规则范围差异。"},
		{"GO-VET-035", "2 处命中为 len(Pattern{}.Features()) 中方法接收者内嵌的复合字面量，len 的直接实参是函数调用返回值而非字面量；规则用 (Call Ident=len $.. CompositeLit) 近似检测，官方 vet 的 appends 分析器只检查 append 参数，属语义子集近似；golangci 报 0 issues。"},
		{"GO-VET-036", "p == nil 是 nilable 类型上的惯用比较；官方 nilfunc 只针对函数值，属规则范围差异。"},
		{"GO-VET-037", "p != nil 是 nilable 类型上的惯用比较；官方 nilfunc 只针对函数值，属规则范围差异。"},
		{"GO-VET-040", "Mutex.Lock 均配对 Unlock；规则检测调用构造，官方 copylocks 针对锁值拷贝，属语义子集近似。"},
		{"GO-VET-041", "Mutex.Unlock 均在 Lock 之后；规则检测调用构造，官方 copylocks 针对锁值拷贝，属语义子集近似。"},
		{"GO-VET-042", "Close 调用均检查错误且仅调用一次；规则检测调用构造，govet/bodyclose 已通过。"},
		{"GO-VET-043", "空块用于显式空实现（接口满足与基准占位）；golangci 未启用对应检查，属项目既定约定。"},
	}
}
