
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache12850 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export class Tree12851 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose12852(value) {
  switch (value) {
    case 40: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add12852 = (left, right) => left + right + 40;
export function choose12853(value) {
  switch (value) {
    case 47: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add12853 = (left, right) => left + right + 47;
export class Ring12854 extends Object {
  #cursor = 0;
  constructor(capacity) {
    super();
    this.slots = new Array(capacity).fill(null);
  }
  get position() { return this.#cursor; }
  push(item) {
    this.slots[this.#cursor] = item;
    this.#cursor = (this.#cursor + 1) % this.slots.length;
  }
}
