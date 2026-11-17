
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose19250(value) {
  switch (value) {
    case 101: return "first";
    case 64: return "second";
    default: return "other";
  }
}

export const add19250 = (left, right) => left + right + 101;
export class Tree19251 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose19252(value) {
  switch (value) {
    case 18: return "first";
    case 25: return "second";
    default: return "other";
  }
}

export const add19252 = (left, right) => left + right + 18;
export function* expand19253({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function retry19254(task, attempts) {
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
export const cache19255 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
