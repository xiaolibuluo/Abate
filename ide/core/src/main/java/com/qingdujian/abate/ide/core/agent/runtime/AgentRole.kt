package com.qingdujian.abate.ide.core.agent.runtime

enum class AgentRole(val displayName: String, val department: String) {
    CEO("CEO Agent", "决策层"),
    CTO("CTO Agent", "决策层"),
    CPO("CPO Agent", "决策层"),
    COO("COO Agent", "决策层"),
    ARCHITECT("架构师 Agent", "工程部"),
    FRONTEND("前端 Agent", "工程部"),
    LOGIC("逻辑 Agent", "工程部"),
    DATA("数据 Agent", "工程部"),
    MODEL("模型 Agent", "工程部"),
    BUILD("构建 Agent", "工程部"),
    UX("UX Agent", "设计部"),
    UI("UI Agent", "设计部"),
    PROTOTYPE("原型 Agent", "设计部"),
    TEST_DESIGN("测试设计 Agent", "QA部"),
    TEST_EXEC("测试执行 Agent", "QA部"),
    CHAOS("混沌 Agent", "QA部"),
    PERFORMANCE("性能 Agent", "QA部"),
    AUDIT("审计 Agent", "安全部"),
    SANDBOX("沙箱 Agent", "安全部"),
    PRIVACY("隐私 Agent", "安全部"),
    CICD("CI/CD Agent", "运维部"),
    RELEASE("发布 Agent", "运维部"),
    MONITOR("监控 Agent", "运维部");

    companion object {
        fun all(): List<AgentRole> = entries.toList()
        fun byDepartment(dept: String): List<AgentRole> = entries.filter { it.department == dept }
    }
}
