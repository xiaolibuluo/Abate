# Abate —— 止息混沌 · Makefile
# Copyright (C) 2026 小黎的部落（www.qingdujian.com）

GO ?= go
GOFLAGS ?=
JAR ?= java -jar tla/tla2tools.jar

.PHONY: build test vet cover tla-check clean

## build: 编译 abate CLI
build:
	$(GO) build $(GOFLAGS) -o abate ./cmd/abate

## test: 运行全部测试
test:
	$(GO) test $(GOFLAGS) ./...

## vet: 静态检查
vet:
	$(GO) vet ./...

## cover: 生成覆盖率报告
cover:
	$(GO) test $(GOFLAGS) -coverprofile=coverage.out ./...
	$(GO) tool cover -func=coverage.out

## tla-check: 运行 TLA+ 模型检验
tla-check:
	cd tla && $(JAR) -config ces.cfg ces.tla

## clean: 清理产物
clean:
	rm -f abate coverage.out
