# Abate 项目统一构建与门禁入口
.PHONY: gate lint vet test build fmt header style license coverage watermark

GOLANGCI := $(shell go env GOPATH)/bin/golangci-lint

# 构建层水印：将构建时间与 git commit 注入 internal/watermark 的包级变量。
WATERMARK_PKG := abate/internal/watermark
BUILD_TIME ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || printf unknown)
WATERMARK_LDFLAGS := -X $(WATERMARK_PKG).BuildTime=$(BUILD_TIME) -X $(WATERMARK_PKG).GitCommit=$(GIT_COMMIT)

# 完整 CI 门禁：任何一环失败即非零退出
gate:
	@bash scripts/lint-gate.sh

# 单独跑全量 linter
lint:
	$(GOLANGCI) run --issues-exit-code=1 --timeout=10m ./...

# 版权头门禁
header:
	@bash scripts/header-check.sh

# 代码风格门禁（300 行 / 一类一个文件）
style:
	@bash scripts/style-check.sh

# 许可证门禁（禁止 GPL）
license:
	@bash scripts/license-check.sh

# 官方 vet 全分析器
vet:
	go vet ./...

# 测试 + 竞态
test:
	go test -race -count=1 ./...

# 覆盖率门禁
coverage:
	@bash scripts/coverage-gate.sh

# 编译（注入构建层水印元数据）
build:
	go build -ldflags "$(WATERMARK_LDFLAGS)" ./...

# 打印并校验三层水印（专用入口）
watermark:
	@go run -ldflags "$(WATERMARK_LDFLAGS)" ./cmd/watermark

# 格式化（提交前必须执行）
fmt:
	gofmt -w .
	$(GOLANGCI) run --fix ./...
