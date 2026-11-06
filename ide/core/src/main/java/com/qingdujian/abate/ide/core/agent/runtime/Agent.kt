package com.qingdujian.abate.ide.core.agent.runtime

import kotlinx.coroutines.flow.Flow

interface Agent {
    val role: AgentRole
    val capability: AgentCapability
    suspend fun setup(): SetupResult
    suspend fun run(task: AgentTask): AgentResult
    suspend fun teardown(): TeardownResult
}

data class AgentCapability(
    val name: String,
    val version: String,
    val protocol: String,
    val tools: List<String>
)

data class AgentTask(
    val id: String,
    val description: String,
    val input: String,
    val dependencies: List<String> = emptyList()
)

data class AgentResult(
    val taskId: String,
    val status: AgentStatus,
    val output: String,
    val metrics: Map<String, Double> = emptyMap()
)

enum class AgentStatus { PASS, FAIL, SKIP, RUNNING }

data class SetupResult(val success: Boolean, val cleanup: (() -> Unit)? = null)
data class TeardownResult(val success: Boolean, val error: String? = null)
