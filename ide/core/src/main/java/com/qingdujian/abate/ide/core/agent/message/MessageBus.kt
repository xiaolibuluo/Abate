package com.qingdujian.abate.ide.core.agent.message

import kotlinx.coroutines.channels.Channel
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.receiveAsFlow
import kotlinx.serialization.Serializable

@Serializable
enum class MessageType { TASK_ASSIGN, TASK_RESULT, TASK_ERROR, TASK_CANCEL }

@Serializable
data class AgentMessage(
    val id: String,
    val type: MessageType,
    val from: String,
    val to: String,
    val payload: String,
    val timestamp: Long
)

class MessageBus {
    private val channel = Channel<AgentMessage>(Channel.UNLIMITED)

    fun send(message: AgentMessage) {
        channel.trySend(message)
    }

    fun receive(): Flow<AgentMessage> = channel.receiveAsFlow()

    fun close() {
        channel.close()
    }
}
