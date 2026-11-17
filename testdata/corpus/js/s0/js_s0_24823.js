
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Tree21870 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function* expand21871({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose21872(value) {
  switch (value) {
    case 29: return "first";
    case 29: return "second";
    default: return "other";
  }
}

export const add21872 = (left, right) => left + right + 29;
export function choose21873(value) {
  switch (value) {
    case 36: return "first";
    case 40: return "second";
    default: return "other";
  }
}

export const add21873 = (left, right) => left + right + 36;
export function choose21874(value) {
  switch (value) {
    case 43: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add21874 = (left, right) => left + right + 43;
export class Tree21875 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export const cache21876 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export const cache21877 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export async function collect21878(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
