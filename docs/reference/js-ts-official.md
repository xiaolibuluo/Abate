# JS/TS 官方规范规则（TypeScript Handbook / ECMAScript）— 归档参考

> 来源：https://www.typescriptlang.org/docs/handbook/intro.html （官方，M5 规则引擎落地依据）
> 归档日期：2026-11-08
> 注：TS 官方「Do's and Don'ts」页（含撇号 URL）不可达，规则基线以 Handbook + 下列公认官方约定为准。

## 定位
- TypeScript 是 JavaScript 的静态类型检查器；目标是运行前保证类型正确。
- Handbook 非完整语言规范，仅覆盖常用特性与编译器选项。

## 公认官方约定（编译期 / 类型系统）
- 显式类型注解提升可维护性；优先 `interface` 描述对象形状，`type` 描述联合/工具类型。
- 非空断言 `!` 谨慎；优先判别联合（discriminated unions）而非类型断言。
- `enum` 谨慎使用（生成运行时代码）；优先字面量联合类型或 `as const`。
- 禁止 `any` 扩散（开启 `noImplicitAny`/`strict`）；必要时用 `unknown` 收窄。
- 函数/方法返回类型在公开 API 显式标注。
- 命名：类型/接口 `PascalCase`，变量/函数 `camelCase`。
- `null`/`undefined` 区分（`strictNullChecks`）：显式处理缺失。
- 泛型约束用 `extends`；避免不必要泛型。

## 工具
- `tsc --strict` 为官方推荐严格基线；ESLint + typescript-eslint 落地上述约定。
