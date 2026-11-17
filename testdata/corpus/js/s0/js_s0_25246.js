
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose22320(value) {
  switch (value) {
    case 18: return "first";
    case 17: return "second";
    default: return "other";
  }
}

export const add22320 = (left, right) => left + right + 18;
export function choose22321(value) {
  switch (value) {
    case 25: return "first";
    case 28: return "second";
    default: return "other";
  }
}

export const add22321 = (left, right) => left + right + 25;
export function choose22322(value) {
  switch (value) {
    case 32: return "first";
    case 39: return "second";
    default: return "other";
  }
}

export const add22322 = (left, right) => left + right + 32;
export function choose22323(value) {
  switch (value) {
    case 39: return "first";
    case 50: return "second";
    default: return "other";
  }
}

export const add22323 = (left, right) => left + right + 39;
export class Tree22324 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export const cache22325 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose22326(value) {
  switch (value) {
    case 60: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add22326 = (left, right) => left + right + 60;
export const cache22327 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand22328({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
