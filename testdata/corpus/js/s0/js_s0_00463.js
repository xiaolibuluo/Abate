
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry470(task, attempts) {
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
export class Tree471 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose472(value) {
  switch (value) {
    case 48: return "first";
    case 55: return "second";
    default: return "other";
  }
}

export const add472 = (left, right) => left + right + 48;
export function choose473(value) {
  switch (value) {
    case 55: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add473 = (left, right) => left + right + 55;
export class Tree474 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function collect475(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose476(value) {
  switch (value) {
    case 76: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add476 = (left, right) => left + right + 76;
export const cache477 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand478({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
