/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// jsRecipe covers ES modules, classes, async/Promise flows, destructuring,
// optional chaining, generators, arrow functions and IIFE closures.
func jsRecipe() Recipe {
	return Recipe{
		Language: "js",
		Ext:      "js",
		Tags:     []string{"class", "async", "destructuring", "generator", "closure", "controlflow", "module"},
		Header:   "\n",
		AnchorClean: `function anchorProbe() {
  return 1 + 2 == 3;
}

`,
		AnchorFull: `function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

`,
		CleanBlock: "function step@@I@@x@@S@@() { return @@N@@ + @@S@@; }\n",
		Blocks: []BlockTemplate{
			{Text: `export class Ring@@I@@@@S@@ extends Object {
  #cursor = 0;
  constructor(capacity) {
    super();
    this.slots = new Array(capacity).fill(null);
  }
  get position() { return this.#cursor; }
  push(item) {
    this.slots[this.#cursor] = item;
    this.#cursor = (this.#cursor + 1) % this.slots.length;
  }
}
`},
			{Text: `export async function collect@@I@@@@S@@(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
`},
			{Text: `export function* expand@@I@@@@S@@({ head, tail = [] }) {
  yield ` + "`${head}:${tail?.length ?? 0}`" + `;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
`},
			{Text: `export function choose@@I@@@@S@@(value) {
  switch (value) {
    case @@N@@: return "first";
    case @@M@@: return "second";
    default: return "other";
  }
}

export const add@@I@@@@S@@ = (left, right) => left + right + @@N@@;
`},
			{Text: `export const cache@@I@@@@S@@ = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
`},
			{Text: `export class Tree@@I@@@@S@@ {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
`},
			{Text: `export async function retry@@I@@@@S@@(task, attempts) {
  for (let attempt = 0; attempt < attempts; attempt++) {
    try {
      return await task(attempt);
    } catch (failure) {
      if (attempt === attempts - 1) {
        throw failure;
      }
    }
  }
  return null;
}
`},
		},
	}
}
