package com.qingdujian.abate.ide.core.model

import kotlinx.coroutines.flow.Flow

interface ModelProvider {
    val name: String
    val isOnline: Boolean
    suspend fun complete(request: ModelRequest): ModelResponse
    suspend fun stream(request: ModelRequest): Flow<ModelChunk>
    fun estimateCost(request: ModelRequest): Int
}

data class ModelRequest(
    val id: String,
    val prompt: String,
    val systemPrompt: String = "",
    val maxTokens: Int = 4096,
    val temperature: Double = 0.7
)

data class ModelResponse(
    val id: String,
    val content: String,
    val tokensUsed: Int,
    val finishReason: String
)

data class ModelChunk(
    val id: String,
    val delta: String,
    val done: Boolean
)
