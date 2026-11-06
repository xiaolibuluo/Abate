package com.qingdujian.abate.ide.core.agent.registry

import com.qingdujian.abate.ide.core.agent.runtime.Agent
import com.qingdujian.abate.ide.core.agent.runtime.AgentRole
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock

class AgentRegistry {
    private val mutex = Mutex()
    private val agents = mutableMapOf<AgentRole, Agent>()

    suspend fun register(agent: Agent): Boolean = mutex.withLock {
        if (agents.containsKey(agent.role)) {
            false
        } else {
            agents[agent.role] = agent
            true
        }
    }

    suspend fun unregister(role: AgentRole): Boolean = mutex.withLock {
        agents.remove(role) != null
    }

    suspend fun lookup(role: AgentRole): Agent? = mutex.withLock {
        agents[role]
    }

    suspend fun list(): List<AgentRole> = mutex.withLock {
        agents.keys.toList().sortedBy { it.ordinal }
    }

    suspend fun count(): Int = mutex.withLock {
        agents.size
    }
}
