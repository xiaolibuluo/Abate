
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose4760(value) {
  switch (value) {
    case 82: return "first";
    case 47: return "second";
    default: return "other";
  }
}

export const add4760 = (left, right) => left + right + 82;
export const cache4761 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export async function retry4762(task, attempts) {
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
export function* expand4763({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose4764(value) {
  switch (value) {
    case 13: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add4764 = (left, right) => left + right + 13;
export class Tree4765 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export const cache4766 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose4767(value) {
  switch (value) {
    case 34: return "first";
    case 63: return "second";
    default: return "other";
  }
}

export const add4767 = (left, right) => left + right + 34;
