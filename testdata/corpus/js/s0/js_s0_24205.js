
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose21250(value) {
  switch (value) {
    case 82: return "first";
    case 61: return "second";
    default: return "other";
  }
}

export const add21250 = (left, right) => left + right + 82;
export const cache21251 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose21252(value) {
  switch (value) {
    case 96: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add21252 = (left, right) => left + right + 96;
export function* expand21253({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function collect21254(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Tree21255 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose21256(value) {
  switch (value) {
    case 27: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add21256 = (left, right) => left + right + 27;
