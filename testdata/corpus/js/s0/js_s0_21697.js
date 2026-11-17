
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache18920 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export async function retry18921(task, attempts) {
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
export async function retry18922(task, attempts) {
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
export function choose18923(value) {
  switch (value) {
    case 81: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add18923 = (left, right) => left + right + 81;
export function choose18924(value) {
  switch (value) {
    case 88: return "first";
    case 60: return "second";
    default: return "other";
  }
}

export const add18924 = (left, right) => left + right + 88;
export function choose18925(value) {
  switch (value) {
    case 95: return "first";
    case 10: return "second";
    default: return "other";
  }
}

export const add18925 = (left, right) => left + right + 95;
export function* expand18926({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function* expand18927({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Tree18928 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
