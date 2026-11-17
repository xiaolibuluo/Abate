
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand12140({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose12141(value) {
  switch (value) {
    case 80: return "first";
    case 39: return "second";
    default: return "other";
  }
}

export const add12141 = (left, right) => left + right + 80;
export class Tree12142 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export const cache12143 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose12144(value) {
  switch (value) {
    case 101: return "first";
    case 11: return "second";
    default: return "other";
  }
}

export const add12144 = (left, right) => left + right + 101;
export const cache12145 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export async function collect12146(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose12147(value) {
  switch (value) {
    case 25: return "first";
    case 44: return "second";
    default: return "other";
  }
}

export const add12147 = (left, right) => left + right + 25;
export async function retry12148(task, attempts) {
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
