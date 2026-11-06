package com.qingdujian.abate.ide.core.tool.permission

import com.qingdujian.abate.ide.core.agent.runtime.AgentRole
import com.qingdujian.abate.ide.core.tool.registry.ToolCategory

enum class Permission { NONE, READ, WRITE }

class PermissionMatrix {
    private val matrix = mutableMapOf<Pair<AgentRole, ToolCategory>, Permission>()

    init {
        // 决策层：只读
        listOf(AgentRole.CEO, AgentRole.CTO, AgentRole.CPO, AgentRole.COO).forEach { role ->
            ToolCategory.entries.forEach { cat ->
                matrix[role to cat] = Permission.READ
            }
        }

        // 工程部：FILE+BUILD 写，其他读
        listOf(AgentRole.ARCHITECT, AgentRole.FRONTEND, AgentRole.LOGIC, AgentRole.DATA, AgentRole.MODEL).forEach { role ->
            matrix[role to ToolCategory.FILE] = Permission.WRITE
            matrix[role to ToolCategory.BUILD] = Permission.READ
            matrix[role to ToolCategory.TEST] = Permission.READ
            matrix[role to ToolCategory.SECURITY] = Permission.READ
            matrix[role to ToolCategory.DESIGN] = if (role == AgentRole.FRONTEND) Permission.WRITE else Permission.READ
            matrix[role to ToolCategory.OPS] = Permission.READ
            matrix[role to ToolCategory.MODEL] = Permission.WRITE
        }

        // 构建 Agent：BUILD 写
        matrix[AgentRole.BUILD to ToolCategory.FILE] = Permission.WRITE
        matrix[AgentRole.BUILD to ToolCategory.BUILD] = Permission.WRITE
        matrix[AgentRole.BUILD to ToolCategory.OPS] = Permission.READ

        // 设计部：FILE+DESIGN 写
        listOf(AgentRole.UX, AgentRole.UI, AgentRole.PROTOTYPE).forEach { role ->
            matrix[role to ToolCategory.FILE] = Permission.WRITE
            matrix[role to ToolCategory.DESIGN] = Permission.WRITE
            matrix[role to ToolCategory.TEST] = Permission.READ
        }

        // QA 部：TEST 写
        listOf(AgentRole.TEST_DESIGN, AgentRole.TEST_EXEC, AgentRole.PERFORMANCE).forEach { role ->
            matrix[role to ToolCategory.TEST] = Permission.WRITE
            matrix[role to ToolCategory.FILE] = if (role == AgentRole.TEST_DESIGN) Permission.WRITE else Permission.READ
        }
        matrix[AgentRole.CHAOS to ToolCategory.TEST] = Permission.WRITE
        matrix[AgentRole.CHAOS to ToolCategory.FILE] = Permission.READ

        // 安全部：SECURITY 写，FILE 只读
        listOf(AgentRole.AUDIT, AgentRole.SANDBOX, AgentRole.PRIVACY).forEach { role ->
            matrix[role to ToolCategory.SECURITY] = Permission.WRITE
            matrix[role to ToolCategory.FILE] = Permission.READ
        }

        // 运维部：BUILD+OPS 写
        listOf(AgentRole.CICD, AgentRole.RELEASE).forEach { role ->
            matrix[role to ToolCategory.BUILD] = Permission.WRITE
            matrix[role to ToolCategory.OPS] = Permission.WRITE
            matrix[role to ToolCategory.FILE] = Permission.WRITE
        }
        matrix[AgentRole.MONITOR to ToolCategory.OPS] = Permission.WRITE
        matrix[AgentRole.MONITOR to ToolCategory.FILE] = Permission.READ
    }

    fun check(role: AgentRole, category: ToolCategory): Permission {
        return matrix[role to category] ?: Permission.NONE
    }

    fun canUse(role: AgentRole, category: ToolCategory): Boolean {
        return check(role, category) != Permission.NONE
    }

    fun canWrite(role: AgentRole, category: ToolCategory): Boolean {
        return check(role, category) == Permission.WRITE
    }
}
