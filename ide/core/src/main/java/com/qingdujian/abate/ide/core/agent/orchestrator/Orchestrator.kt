package com.qingdujian.abate.ide.core.agent.orchestrator

import com.qingdujian.abate.ide.core.agent.message.AgentMessage
import com.qingdujian.abate.ide.core.agent.message.MessageBus
import com.qingdujian.abate.ide.core.agent.message.MessageType
import com.qingdujian.abate.ide.core.agent.registry.AgentRegistry
import com.qingdujian.abate.ide.core.agent.runtime.*
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope

enum class FailureStrategy { CONTINUE, ABORT_DEPENDENTS, ABORT_ALL }

data class TaskNode(
    val id: String,
    val role: AgentRole,
    val task: AgentTask,
    val dependsOn: List<String> = emptyList()
)

data class OrchestratorReport(
    val total: Int,
    val passed: Int,
    val failed: Int,
    val skipped: Int,
    val results: List<AgentResult>
)

class Orchestrator(
    private val registry: AgentRegistry,
    private val bus: MessageBus,
    private val strategy: FailureStrategy = FailureStrategy.ABORT_DEPENDENTS
) {
    suspend fun execute(nodes: List<TaskNode>): OrchestratorReport = coroutineScope {
        val results = mutableMapOf<String, AgentResult>()
        val completed = mutableSetOf<String>()
        val failed = mutableSetOf<String>()

        while (completed.size + failed.size < nodes.size) {
            val ready = nodes.filter { node ->
                node.id !in completed && node.id !in failed &&
                node.dependsOn.all { it in completed }
            }

            if (ready.isEmpty()) {
                val blocked = nodes.filter { it.id !in completed && it.id !in failed }
                blocked.forEach { node ->
                    results[node.id] = AgentResult(node.id, AgentStatus.SKIP, "blocked by failed dependency")
                    failed.add(node.id)
                }
                break
            }

            val deferred = ready.map { node ->
                async {
                    val agent = registry.lookup(node.role)
                    if (agent == null) {
                        AgentResult(node.id, AgentStatus.SKIP, "agent not registered")
                    } else {
                        try {
                            agent.setup()
                            val result = agent.run(node.task)
                            agent.teardown()
                            result
                        } catch (e: Exception) {
                            AgentResult(node.id, AgentStatus.FAIL, e.message ?: "unknown error")
                        }
                    }
                }
            }

            val batchResults = deferred.awaitAll()
            batchResults.forEach { result ->
                results[result.taskId] = result
                when (result.status) {
                    AgentStatus.PASS -> completed.add(result.taskId)
                    AgentStatus.FAIL -> {
                        failed.add(result.taskId)
                        if (strategy == FailureStrategy.ABORT_ALL) {
                            nodes.filter { it.id !in completed && it.id !in failed }.forEach {
                                results[it.id] = AgentResult(it.id, AgentStatus.SKIP, "aborted")
                                failed.add(it.id)
                            }
                        }
                    }
                    else -> completed.add(result.taskId)
                }
            }
        }

        val allResults = results.values.toList()
        OrchestratorReport(
            total = nodes.size,
            passed = allResults.count { it.status == AgentStatus.PASS },
            failed = allResults.count { it.status == AgentStatus.FAIL },
            skipped = allResults.count { it.status == AgentStatus.SKIP },
            results = allResults
        )
    }
}
