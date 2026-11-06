package com.qingdujian.abate.ide.core.tool.registry

import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock

data class Tool(
    val name: String,
    val description: String,
    val category: ToolCategory
)

enum class ToolCategory {
    FILE, BUILD, TEST, SECURITY, DESIGN, OPS, MODEL
}

class ToolRegistry {
    private val mutex = Mutex()
    private val tools = mutableMapOf<String, Tool>()

    suspend fun register(tool: Tool): Boolean = mutex.withLock {
        if (tools.containsKey(tool.name)) false
        else { tools[tool.name] = tool; true }
    }

    suspend fun lookup(name: String): Tool? = mutex.withLock { tools[name] }

    suspend fun list(): List<Tool> = mutex.withLock { tools.values.toList() }

    suspend fun byCategory(cat: ToolCategory): List<Tool> = mutex.withLock {
        tools.values.filter { it.category == cat }
    }

    companion object {
        fun defaultTools(): List<Tool> = listOf(
            Tool("read_file", "读取文件", ToolCategory.FILE),
            Tool("write_file", "写入文件", ToolCategory.FILE),
            Tool("search_code", "搜索代码", ToolCategory.FILE),
            Tool("glob_files", "通配符查找", ToolCategory.FILE),
            Tool("dependency_tree", "依赖分析", ToolCategory.FILE),
            Tool("build_run", "Gradle 构建", ToolCategory.BUILD),
            Tool("sign_config", "签名配置", ToolCategory.BUILD),
            Tool("variant_manage", "构建变体", ToolCategory.BUILD),
            Tool("cache_manage", "构建缓存", ToolCategory.BUILD),
            Tool("test_run", "运行测试", ToolCategory.TEST),
            Tool("test_report", "测试报告", ToolCategory.TEST),
            Tool("coverage_report", "覆盖率", ToolCategory.TEST),
            Tool("abate_inject", "混沌注入", ToolCategory.TEST),
            Tool("abate_report", "混沌报告", ToolCategory.TEST),
            Tool("profile_run", "性能分析", ToolCategory.TEST),
            Tool("trace_view", "Trace 查看", ToolCategory.TEST),
            Tool("static_scan", "静态安全扫描", ToolCategory.SECURITY),
            Tool("permission_audit", "权限审查", ToolCategory.SECURITY),
            Tool("sandbox_test", "沙箱测试", ToolCategory.SECURITY),
            Tool("isolation_verify", "隔离验证", ToolCategory.SECURITY),
            Tool("compliance_check", "合规检查", ToolCategory.SECURITY),
            Tool("data_flow_audit", "数据流审计", ToolCategory.SECURITY),
            Tool("preview_render", "Compose 预览", ToolCategory.DESIGN),
            Tool("layout_inspector", "布局检查", ToolCategory.DESIGN),
            Tool("flowchart_render", "流程图", ToolCategory.DESIGN),
            Tool("color_palette", "色彩方案", ToolCategory.DESIGN),
            Tool("theme_generate", "主题生成", ToolCategory.DESIGN),
            Tool("spec_write", "规范文档", ToolCategory.DESIGN),
            Tool("pipeline_run", "流水线", ToolCategory.OPS),
            Tool("version_tag", "版本标记", ToolCategory.OPS),
            Tool("release_check", "发布检查", ToolCategory.OPS),
            Tool("crash_report", "崩溃报告", ToolCategory.OPS),
            Tool("anr_detect", "ANR 检测", ToolCategory.OPS)
        )
    }
}
