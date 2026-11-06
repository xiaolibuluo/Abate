# Abate IDE —— 模型层架构

> Copyright (C) 2026 小黎的部落（www.qingdujian.com）
> AGPLv3; 商业用途须事先获得小黎的部落书面授权

## 1. 架构

```
Model Abstraction Layer (统一接口)
  |
  +-- API Models (在线)
  |     OpenAI / Claude / Gemini / 自定义端点
  |     OkHttp + SSE 流式
  |
  +-- Local Models (离线)
  |     ONNX 量化 / GGUF / 用户上传
  |     ONNX Runtime + MLC-LLM
  |
  +-- Model Router (智能路由)
        任务匹配 / 成本控制 / 离线降级 / 质量评估
```

## 2. API 模型适配

| 提供商 | 协议 | 特点 |
|--------|------|------|
| OpenAI | OpenAI API + SSE | GPT 系列，流式输出 |
| Claude | Anthropic API + SSE | 长上下文，代码强 |
| Gemini | Google API + SSE | 多模态 |
| 自定义 | 兼容 OpenAI 格式 | 用户填 URL + Key |

- API Key 加密存储于 Android Keystore
- 支持流式输出（SSE），实时展示 Agent 思考过程
- 自动重试 + 超时控制 + Token 计数

## 3. 本地模型

| 格式 | 推理引擎 | 特点 |
|------|----------|------|
| ONNX | ONNX Runtime | 量化压缩，CPU 推理 |
| GGUF | MLC-LLM | GPU 加速，大模型 |
| 用户上传 | 自动检测格式 | 支持 .onnx/.gguf/.tflite |

- 模型文件存储在应用私有目录
- 加载时限制内存配额（防止 OOM）
- 支持模型卸载释放内存
- 推理在独立线程，不阻塞 UI

## 4. 智能路由策略

| 任务类型 | 优先模型 | 降级 |
|----------|----------|------|
| 架构决策/复杂推理 | API 大模型 | 本地大模型 |
| 代码生成/补全 | 用户配置 | 本地模型 |
| 简单检查/格式化 | 本地小模型 | 无 |
| 安全审计 | API 大模型 | 本地大模型 |
| 混沌测试分析 | 本地模型 | API 模型 |

### 路由决策因子

1. 任务类型（推理复杂度）
2. 网络状态（在线/离线）
3. 成本预算（Token 计数）
4. 用户配置（手动指定优先）
5. 模型负载（排队等待）
6. 质量评估（历史成功率）

## 5. 用户配置

- 每个 Agent 角色可单独绑定模型
- 全局默认模型配置
- 离线模式开关（强制使用本地模型）
- 成本上限设置（每日 Token 预算）
- 模型测试功能（发送测试 prompt 验证连通性）

## 6. 统一接口

```kotlin
interface ModelProvider {
    val name: String
    val isOnline: Boolean
    suspend fun complete(request: ModelRequest): ModelResponse
    suspend fun stream(request: ModelRequest): Flow<ModelChunk>
    fun estimateCost(request: ModelRequest): Int
}
```

所有模型适配器实现统一接口，Agent Runtime 不感知底层差异。