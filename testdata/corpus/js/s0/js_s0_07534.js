
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache6600 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose6601(value) {
  switch (value) {
    case 56: return "first";
    case 65: return "second";
    default: return "other";
  }
}

export const add6601 = (left, right) => left + right + 56;
export function choose6602(value) {
  switch (value) {
    case 63: return "first";
    case 15: return "second";
    default: return "other";
  }
}

export const add6602 = (left, right) => left + right + 63;
export const cache6603 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand6604({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Tree6605 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
