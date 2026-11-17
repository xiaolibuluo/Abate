
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect4240(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export const cache4241 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand4242({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose4243(value) {
  switch (value) {
    case 9: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add4243 = (left, right) => left + right + 9;
export function choose4244(value) {
  switch (value) {
    case 16: return "first";
    case 32: return "second";
    default: return "other";
  }
}

export const add4244 = (left, right) => left + right + 16;
export class Tree4245 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose4246(value) {
  switch (value) {
    case 30: return "first";
    case 54: return "second";
    default: return "other";
  }
}

export const add4246 = (left, right) => left + right + 30;
