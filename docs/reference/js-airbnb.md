# JS/TS 社区规范规则（Airbnb JavaScript Style Guide）— 归档参考

> 来源：https://airbnb.io/javascript/ （社区，M5 规则引擎落地依据）
> 归档日期：2026-11-08

## 变量声明
- 所有引用用 `const`，重赋值用 `let`，杜绝 `var`；每个变量单独声明；禁止链式赋值；避免 `++`/`--`（用 `num += 1`）。

## 对象 / 数组
- 字面量 `{}`/`[]`；方法/属性简写；浅拷贝用展开 `{...orig}`/`[...items]` 而非 `Object.assign`。
- 数组添加用 `push`；复制用展开；回调必须有 `return`。

## 字符串
- 单引号；程序化构建用模板字符串；禁止 `eval()`。

## 函数
- 命名函数表达式（`func-style`）；默认参数语法（置于末尾）；禁止 `arguments`（用 `...args`）；禁止 `new Function()`；禁止修改/重赋值参数。

## 箭头函数
- 匿名回调用箭头；单表达式可省 `{}` 隐式返回；单参数省括号 `x => x`。

## 类 / 模块
- 始终用 `class` + `extends`；始终用 ES Module（`import`/`export`）；禁止通配符导入 `import *`；禁止导出可变绑定；`import` 置于顶部。

## 迭代器
- 用 `map`/`filter`/`reduce` 而非 `for-in`/`for-of`。

## Do / Don't
- **Always**：const/let 非 var；默认参数末位；class 语法；ES Module；字面量；单引号；先 const 后 let 分组。
- **Never**：字符串 eval；参数命名 arguments；`new Function`；修改/重赋值参数；块内声明函数；链式赋值；迭代器/生成器；通配符导入；导出可变绑定。
