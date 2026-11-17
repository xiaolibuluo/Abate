# Abate L7 控制面规划 — 三模块极致细节

- 文档编号：PLAN-07　状态：生效　上游：PLAN-00
- 覆盖模块：CTL-CLI / CTL-REST / CTL-WEB
- 通用约束（三模块共同遵守）：
  1. 控制面是纯编排层：只调用 L0-L6 暴露的端口接口，禁止内嵌任何业务逻辑（检查规则、算子语义、验证判定一律不得出现在 L7）。
  2. 控制面自身零状态：实验状态全部在内核（调度器/档案）；CLI 每次运行独立，REST/Web 经内核会话接口操作。
  3. 退出码约定（三入口统一）：0 = 成功；1 = 门禁/实验判定失败（业务失败）；2 = 内部错误；3 = 用法错误（参数/配置非法）。业务失败与内部错误严格区分（CI 消费语义）。
  4. 覆盖率：语句 ≥85%、分支 ≥75%（生态层标准）。

---

## 模块 CTL-CLI：命令行

### 1. 模块定位与职责

abate 命令行：用户与 CI 的主入口。职责：命令解析与分发、配置加载、实验驱动（run/check/mutate/chaos）、档案管理（replay/verify）、插件管理（install/list/upgrade）、报告输出。M1 交付最小集（run/check），随里程碑逐步补全。

### 2. 包结构与文件布局

```
ctl/cli/
  main.go          // 入口：构建根命令并 Execute
  root.go          // 根命令：全局 flag（--config/--verbose/--no-color）
  cmd_run.go       // run：执行 CES 实验
  cmd_check.go     // check：语言代码检查
  cmd_mutate.go    // mutate：变异实验
  cmd_chaos.go     // chaos：混沌实验（run 的语义快捷方式）
  cmd_replay.go    // replay：档案回放
  cmd_verify.go    // verify：双跑比对（确定性验收）
  cmd_plugin.go    // plugin install/list/upgrade/remove
  cmd_report.go    // report：从档案生成报告（格式转换）
  cmd_version.go   // version：版本与构建元数据
  config.go        // 配置加载（默认路径查找 + 合并优先级）
  exit.go          // 退出码映射
  cli_test.go
```

### 3. 命令树全枚举（首版必须齐备）

```
abate
├── run <ces-file>              # 执行 CES 实验
│   --seed uint                  # 确定性种子（缺省生成并回写报告）
│   --format sarif|json|junit|console|html
│   --out path                   # 报告输出路径
│   --record path                # 档案录制路径
│   --budget-events/-duration/-cost  # 预算覆盖
│   --dry-run                    # 只解析+校验，不执行
│   --strict                     # warning 违规也 FAIL（默认仅 error）
├── check <paths...>             # 语言代码检查
│   --lang go|python|js|c|rust|csharp|kotlin|all
│   --rules id-filter            # 规则过滤（逗号分隔/前缀通配）
│   --severity error|warning|info  # 阈值
│   --format sarif|json|console
│   --exit-zero                  # 违规不置 FAIL（报告模式）
├── mutate <paths...>            # 变异实验
│   --suite test-selector        # 测试套件选择器
│   --operators id-filter        # 算子过滤
│   --max-mutants int             # 变异体上限
│   --format sarif|json|junit|html
├── chaos <ces-file>              # 混沌实验（= run 的 C 族快捷方式）
├── replay <archive>              # 档案回放
│   --verify                     # 回放中逐帧校验
├── verify <a> <b>               # 双跑比对（确定性验收）
│   --diff-out path              # 差异报告输出
├── plugin
│   ├── install <path|ref>       # 安装插件（验签→能力审查→入库）
│   ├── list                     # 已装插件清单（类型/版本/能力/签名状态）
│   ├── upgrade [id]             # 升级（热加载）
│   └── remove <id>              # 移除
├── report <archive|json-report>  # 从档案/JSON 报告生成其他格式
│   --format html|sarif|junit
├── version                      # 版本/commit/构建元数据/水印
└── gate                         # 完整门禁链（scripts/lint-gate.sh 的内置等价物）
```

### 4. 核心类型与接口（必须实现）

- 根命令全局 flag：`--config`（配置文件路径；默认查找顺序：./abate.yml → ~/.config/abate/abate.yml → 内置默认）、`--verbose`（日志级别）、`--no-color`（CI 环境自动检测 TTY 关色）。
- 配置模型（abate.yml）：plugins（安装目录/信任锚附加路径）、report（默认格式/输出目录）、budget（默认预算三上限）、net（离线模式开关——默认离线，铁律）。
- 退出码映射：run/check/mutate/chaos/gate → 0/1/2/3 四态；plugin/report/replay/verify → 0/2/3（无业务失败态，verify 差异返回 1）。
- run 执行链：加载配置 → CES-PARSE（失败 exit 3 + 结构化错误）→ CES-VALID（error 违规 exit 3；--strict 时 warning 也 exit 3）→ 调度执行（KRN-SCH）→ 报告输出（RPT-*）→ 判定映射退出码（稳态 VIOLATED 或变异 SURVIVED 超阈值 → exit 1）。
- verify 执行链：加载两档案 → SIM-REC Compare → 零差异 exit 0；有差异输出结构化 Diff（首个差异点 + 统计）exit 1。

### 5. 关键算法与流程

1. 命令分发：cobra 风格子命令注册（源码级并入 spf13/cobra？否——标准库 flag 自研轻量命令树，避免 vendor；ADR 记录决策：L7 允许第三方源码并入，但 CLI 框架自研以控依赖面）。
2. 帮助系统：每命令 --help 输出用法/示例/退出码语义；示例全部可复制执行（文档测试校验）。
3. 进度输出：长任务（mutate 大规模变异体）进度条（--verbose 时逐条输出；默认单行刷新；--no-color 纯文本百分比）。
4. 信号处理：SIGINT/SIGTERM 触发优雅终止（调度器停止 → 档案落盘 → 报告部分输出标记 incomplete）——禁止直接 exit 丢档案。

### 6. 依赖与约束

- 依赖：L0-L6 全部端口接口（经 use-case 编排层）、标准库。
- 禁止：CLI 内嵌业务判定；网络（除 plugin install 显式网络操作且 --offline 拒绝）；第三方 CLI 框架 vendor（自研命令树）。

### 7. 测试要求

- 单元：命令树解析（每命令 ≥3 用例：合法/缺参/非法参）；退出码四态映射；配置查找合并优先级。
- 端到端（CLI 级集成）：run（合法/非法 CES 各用例，断言退出码与输出）；check（夹具工程断言违规清单）；verify（同种子双跑零差异 + 异种子差异检出）；plugin install（未签名拒绝）。
- 文档测试：--help 输出的示例命令全部实跑通过。
- 信号测试：SIGINT 中断后档案完整落盘（incomplete 标记）。

### 8. 验收标准（可量化）

1. 命令树全枚举实现齐备；每命令帮助与示例文档测试 100% 通过。
2. 退出码四态映射用例 100% 通过（CI 消费语义正确）。
3. 端到端用例（run/check/verify/plugin ≥12 条）100% 通过。
4. 信号优雅终止用例通过（档案不丢失）。
5. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 CTL-REST：REST API

### 1. 模块定位与职责

本地/远程控制 API：实验提交与查询、报告获取、插件管理、健康探针。服务长驻场景（Web 控制台后端、远程编排系统对接）。默认仅监听 localhost（远程监听需显式配置 + 强制鉴权）。

### 2. 包结构与文件布局

```
ctl/rest/
  server.go       // Server：路由注册与生命周期
  auth.go         // 鉴权（token；远程监听强制）
  middleware.go   // 中间件链：日志/限流/恢复/超时
  api_experiment.go // 实验资源端点
  api_report.go    // 报告资源端点
  api_plugin.go    // 插件资源端点
  api_health.go    // 健康探针
  openapi.go      // OpenAPI 3.1 文档生成
  rest_test.go
```

### 3. REST 资源与端点全定义（首版必须齐备）

- 实验：
  - POST /api/v1/experiments：提交 CES 文档（body：YAML/JSON）→ 202 + Location: /experiments/{id}（异步执行）。
  - GET /api/v1/experiments/{id}：状态（pending/running/completed/failed/incomplete）+ 摘要（分数/判定）。
  - GET /api/v1/experiments/{id}/events：事件流（SSE，实时进度）。
  - POST /api/v1/experiments/{id}/cancel：取消（优雅终止）。
  - GET /api/v1/experiments：列表（分页/过滤：状态/时间窗）。
- 报告：
  - GET /api/v1/experiments/{id}/report?format=sarif|json|junit|html：报告获取。
  - GET /api/v1/experiments/{id}/archive：档案下载（SIM-REC 二进制）。
- 检查：
  - POST /api/v1/checks：提交检查任务（paths + 选项）→ 202 异步 + 同 CLI check 语义。
  - GET /api/v1/checks/{id}：结果（违规清单 + 摘要）。
- 插件：
  - GET /api/v1/plugins：清单（类型/版本/能力/签名状态）。
  - POST /api/v1/plugins：安装（multipart 上传 .wasm+.sig → 验签 → 入库）。
  - DELETE /api/v1/plugins/{id}：移除。
  - POST /api/v1/plugins/{id}/upgrade：升级（热加载）。
- 系统：
  - GET /api/v1/health：存活（liveness：进程可响应）。
  - GET /api/v1/ready：就绪（readiness：内核初始化完成 + 信任锚加载）。
  - GET /api/v1/version：版本/构建元数据。
  - GET /openapi.json：OpenAPI 3.1 文档（生成物）。

### 4. 核心类型与接口（必须实现）

- `Server`：`New(cfg ServerConfig) (*Server, error)`、`ListenAndServe() error`、`Shutdown(ctx) error`（优雅：排空在途请求 → 终止运行中实验 → 落档案）。
- ServerConfig：Listen（默认 127.0.0.1:8420）、Auth（Token；远程监听 0.0.0.0 时强制非空，否则启动拒绝）、MaxConcurrent（并发实验上限，默认 CPU 数）、RequestTimeout（默认 30s；长操作一律异步任务化）。
- 鉴权：Authorization: Bearer <token>；除 /health /ready /version 外全部端点强制（本地监听同样强制——防本机恶意进程；文档明示）。
- 任务模型：异步执行（202 + 轮询/SSE）；任务队列上限 100（满则 429 + Retry-After）。
- 错误模型（RFC 9457 problem+json）：type/title/status/detail/instance；错误码与 CLI 退出码语义对齐（validation→400、not-found→404、conflict→409、internal→500）。

### 5. 关键算法与流程

1. 中间件链（固定顺序）：恢复（panic→500+日志）→ 日志（结构化：方法/路径/状态/耗时/请求 ID）→ 超时 → 限流（令牌桶：100 req/s/IP 默认）→ 鉴权 → 路由。
2. SSE 事件流：实验事件（状态转移/进度）经订阅 SIM-REC 帧流转发；断线重连（Last-Event-ID 续传，档案 Seq 为游标）。
3. OpenAPI 生成：路由注册时声明元数据（参数/响应 schema）→ 运行时聚合生成 openapi.json；CI 校验生成物与路由零漂移。
4. 并发安全：实验任务互隔离（各自调度器实例）；插件操作串行化（单写锁——热加载原子性）。

### 6. 依赖与约束

- 依赖：L0-L6 端口接口、标准库 net/http（Go 1.22+ 路由增强，不引第三方路由框架——ADR 记录）。
- 禁止：业务逻辑内嵌；远程监听无鉴权（启动拒绝）；明文 token 入日志（脱敏）。

### 7. 测试要求

- 单元：路由表全端点（每端点 ≥2 用例：合法/非法）；鉴权（无 token/错 token 401）；限流（超限 429）；错误模型（problem+json 结构断言）。
- 集成：httptest 全链路（提交实验→轮询完成→取报告→取档案）；SSE（断线重连续传）；取消（cancel→incomplete 状态+档案完整）。
- 安全：远程监听无 token 启动拒绝用例；token 日志脱敏断言。
- OpenAPI：生成物 schema 校验 + 路由零漂移 CI 检查。

### 8. 验收标准（可量化）

1. 全端点（实验 5 + 报告 2 + 检查 2 + 插件 4 + 系统 4 = 17 端点）用例 100% 通过。
2. OpenAPI 3.1 生成物过官方校验器；路由零漂移 CI 检查通过。
3. SSE 断线重连续传用例通过；取消优雅终止用例通过（档案完整）。
4. 安全用例（强制鉴权/脱敏/远程拒绝）100% 通过。
5. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 CTL-WEB：Web 控制台

### 1. 模块定位与职责

浏览器控制台：实验提交与监控、报告浏览、变异矩阵交互、插件管理。纯静态前端（SPA）+ CTL-REST 后端；前端资源由 abate web 命令静态服务（同源，无 CORS 面）。

### 2. 包结构与文件布局

```
ctl/web/
  serve.go        // abate web 命令：起 REST + 静态服务
  embed.go        // go:embed 前端构建产物（dist/）
  dist/           // 前端构建产物（入库，保证离线可用）
  web_test.go
frontend/         // 前端源码（独立子项目，构建产物输出 ctl/web/dist/）
  src/
    pages/        # 页面组件
    api/          # REST 客户端（OpenAPI 生成）
    components/   # 通用组件
```

### 3. 页面清单（首版必须齐备）

1. 实验列表页：状态过滤/分页；行内摘要（分数/判定/时长）。
2. 实验详情页：HPOV 状态时间线（状态流转可视化）+ 实时事件流（SSE）+ 取消按钮。
3. 实验提交页：CES 文档编辑器（语法高亮/校验错误内联标注——复用 CES-PARSE 错误位置）+ 种子/预算表单。
4. 报告页：违规清单（筛选/排序/源码片段±3 行）+ 变异矩阵（行=变异体，列=测试，击杀标记）+ 分数仪表。
5. 插件页：清单（类型/版本/能力/签名状态）+ 安装（上传 .wasm+.sig）+ 升级/移除。
6. 设置页：服务地址/token 配置（存浏览器本地）。

### 4. 核心类型与接口（必须实现）

- `abate web` 命令：`--listen`（默认 127.0.0.1:8430）+ `--api`（REST 地址，默认同进程内嵌启动）；静态服务（go:embed dist；SPA 路由回退 index.html）。
- REST 客户端：OpenAPI 生成（类型安全；生成物入库；CI 校验与 openapi.json 零漂移）。
- 前端技术约束：无运行时框架 CDN 依赖（构建期打包，产物自包含）；构建工具链锁定版本（pnpm-lock 入库）。

### 5. 关键算法与流程

1. 同源服务：web 命令内嵌 REST（默认）或反代指定 API（--api）；同源消除 CORS 配置面。
2. SSE 消费：实验详情页订阅事件流；断线自动重连（Last-Event-ID）+ 状态提示。
3. 报告渲染：SARIF/JSON 数据驱动渲染；大变异矩阵虚拟滚动（>1000 行不卡顿）。

### 6. 依赖与约束

- 依赖：CTL-REST（API）、前端构建工具链（构建期，不入运行时）。
- 禁止：运行时外部 CDN/字体/脚本（完全自包含）；前端内嵌业务逻辑（纯展示与编排，判定全部来自 API）。

### 7. 测试要求

- 前端单测：页面组件（每页 ≥3 用例：数据渲染/空态/错误态）；API 客户端生成物类型检查。
- E2E：Playwright（源码级并入工具链，dev 依赖）六页面全流程（提交→监控→报告→插件管理）。
- 构建：dist 产物入库与构建可复现（CI 双构建字节一致——前端确定性）。

### 8. 验收标准（可量化）

1. 六页面组件单测 100% 通过；E2E 全流程通过。
2. API 客户端生成物与 openapi.json 零漂移（CI 检查）。
3. dist 构建可复现（双构建字节一致）；产物自包含（无外部请求——网络面板断言）。
4. 大矩阵虚拟滚动性能：1 万变异体渲染交互 P99 < 100ms（E2E 性能断言）。
5. 覆盖率 ≥85%/75%（Go 侧）；门禁全绿。

---

## 附：L7 合入总验收（三模块齐备后）

1. 端到端：CLI 提交实验 → REST 查询状态 → Web 监控 → 报告三端一致（CLI 输出/REST 报告/Web 渲染同源数据断言）。
2. 退出码/API 错误码/OpenAPI 三方语义一致（映射表入库，CI 校验）。
3. 离线验证：断网环境 CLI 全命令 + Web 全页面可用（零外部依赖铁律）。
4. 三模块覆盖率汇总报告入库；五段门禁+基准回归全绿。
