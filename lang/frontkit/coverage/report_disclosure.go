/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

// degradeFacts 逐语言登记审计报告第 7 维度的降级率实测值。
var degradeFacts = map[string]string{
	"c":  "无冻结降级率阈值；全语料 **28.00%**（超阈）。",
	"cs": "单文件 **3.26%**（≤5 达标）；全语料 **6.98%**（超阈）。",
	"js": "无冻结降级率阈值；全语料 **49.12%**（超阈）。",
	"kt": "单文件 **4.48%**（≤5 达标）；全语料 **8.21%**（超阈）。",
	"py": "单文件 **0%**（≤3 达标）；全语料 **26.89%**（超阈）。",
	"rs": "单文件 **0.61%**（≤3 达标）；全语料 **42.33%**（超阈）。",
}

// degradeDisclosure 返回一份报告的降级率口径披露段。
func degradeDisclosure(lang string) string {
	values, ok := degradeFacts[lang]
	if !ok {
		values = "无该语言的披露数据，如实标记为 NOT_MEASURED。"
	}
	return "## 降级率口径说明（诚实披露）\n\n" +
		"`degrade*Pct` 按每语言**首个真实解析 unit（单文件）**计算，" +
		"并对应冻结阈值；全语料使用同一 `Degradation` 公式，" +
		"数值不同——两者公式相同、仅统计范围不同。\n\n" +
		"- 本语言：" + values + "\n" +
		"- 结论限定：达标仅限冻结的单文件口径，" +
		"不代表全语料/语料级降级率达标；go 为 **NOT_MEASURED**。\n" +
		"- 数据来源：`docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 第 7 维度；" +
		"数值真实、口径不同，禁止把单文件达标外推为全语料达标。\n\n"
}
