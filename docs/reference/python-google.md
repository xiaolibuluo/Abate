# Python 社区规范规则（Google Python Style Guide）— 归档参考

> 来源：https://google.github.io/styleguide/pyguide.html （社区，M5 规则引擎落地依据）
> 归档日期：2026-11-08

## 行长度
- ≤80 字符；禁止反斜杠换行，用括号隐式换行；自动格式化（Black/Pyink）后可豁免。

## 导入
- 仅 `import` 包/模块，禁止导入单个类/函数（typing 等豁免）；完整包名禁止相对导入；分组：future→标准库→第三方→仓库子包，组内字典序。

## 格式化
- 4 空格，禁 tab；少括号；顶层定义间 2 空行、方法间 1 空行。

## 命名（遵循 PEP 8）
- 模块 `lower_with_under.py`；类 `CapWords`；函数 `lower_with_under()`；常量 `CAPS_WITH_UNDER`；异常 `CapWords`+`Error`。

## Docstring
- 三双引号（PEP 257）；摘要 ≤80 字符；函数需 `Args/Returns/Yields/Raises`。

## 类型注解（PEP 484）
- 公开 API 强烈建议注解；`X | None` 或 `Optional`；禁止隐式 `a: str = None`；变量注解禁行尾 `# type:` 注释。

## 错误处理
- 用内置异常类；禁止 `except:` / `except Exception` 泛捕获（除非重抛/隔离记录）；`try` 块最小化；`None` 用 `is None`；资源用 `with`。

## 工具
- 必须 Pylint；推荐 Black/Pyink 自动格式化；类型检查 pyrefly。
