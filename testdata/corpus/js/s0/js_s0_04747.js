
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache4120 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export class Tree4121 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose4122(value) {
  switch (value) {
    case 40: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add4122 = (left, right) => left + right + 40;
export const cache4123 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose4124(value) {
  switch (value) {
    case 54: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add4124 = (left, right) => left + right + 54;
