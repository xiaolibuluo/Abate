
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache14000 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose14001(value) {
  switch (value) {
    case 73: return "first";
    case 60: return "second";
    default: return "other";
  }
}

export const add14001 = (left, right) => left + right + 73;
export function* expand14002({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose14003(value) {
  switch (value) {
    case 87: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add14003 = (left, right) => left + right + 87;
export class Tree14004 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export const cache14005 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose14006(value) {
  switch (value) {
    case 11: return "first";
    case 54: return "second";
    default: return "other";
  }
}

export const add14006 = (left, right) => left + right + 11;
export class Tree14007 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
