# Kotlin 官方规范规则（Kotlin Coding Conventions）— 归档参考

> 来源：https://kotlinlang.org/docs/coding-conventions.html （官方，M5 规则引擎落地依据）
> 归档日期：2026-11-08

## 命名
- 类/对象 `UpperCamelCase`；函数 camelCase 无下划线；工厂函数可同名抽象返回类型；测试函数允许反引号（`ensure everything works`）；两字母缩写全大写（`IOStream`），超两字母仅首字母大写（`XmlFormatter`）。
- 属性/局部变量 camelCase；backing property 私有 `_elementList` + 公开 `elementList`。
- 常量（`const` 或顶层不可变 `val`）：`MAX_COUNT` 全大写下划线；枚举常量全大写或 UpperCamelCase 均可。

## 格式化
- 4 空格缩进，禁 tab。
- 左大括号行尾；右大括号单独成行对齐。
- 二元运算符两侧空格（`0..i` 范围运算符除外）；一元运算符无空格；控制流关键字与 `(` 间空格；`//` 后空格；冒号 `:` 分隔类型与父类型前加空格、声明与类型间不加、冒号后始终加空格。

## 源文件组织
- 单类用 `ClassName.kt`；多声明用描述性 UpperCamelCase（禁 `Util`）；多平台源集加后缀（`Platform.jvm.kt`）。
- 类内部顺序：属性+初始化块 → 次构造器 → 方法 → 伴生对象；重载方法相邻；接口实现顺序与接口一致。

## Lambda
- `{ it > 10 }` 花括号与箭头周围空格；单参数尽量移出括号；嵌套 lambda 必须显式声明参数。

## Companion Objects
- 必须位于类最末位；工厂函数可定义在 companion 中（建议区分名如 `fromPolar`）。
