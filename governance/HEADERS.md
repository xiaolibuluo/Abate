<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 版权头规范与水印（XCU-GOV）

> 对应 PLAN-08 XCU-GOV §3。所有源文件必须携带统一版权头；门禁由
> `scripts/header-check.sh` 常驻校验。

## 1. Go 源文件模板（逐字节复用）

```go
/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
```

`main.go` 前 7 行（含空行）为权威样式，新增 `.go` 文件须逐字节复用此头。

## 2. 覆盖范围

- `.go` / `.sh` / `.yml` / `.md`（文档头部 HTML 注释或 YAML front-matter）。
- `third_party/` **排除**：保留上游版权头，由 `third_party/LICENSE-MANIFEST.md` 登记管理。

## 3. 门禁规则

`scripts/header-check.sh` 取文件前 8 行，缺失以下任一标记即 FAIL：

- `Copyright`
- `xiaolibuluo/Abate`（仓库来源）
- `小黎的部落`（数字水印）

## 4. 数字代码水印（三重）

1. **版权头水印**：上述模板（静态，随源码分发）。
2. **构建元数据水印**：`-ldflags` 注入版本/commit/构建时间/水印标识『小黎的部落』；
   `abate version` 输出全量元数据（见 `scripts/release.sh`）。
3. **发布清单水印**：`manifest.txt` 含水印行；GPG 签名覆盖（防篡改）。
