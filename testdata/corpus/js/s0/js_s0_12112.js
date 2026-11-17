
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose10910(value) {
  switch (value) {
    case 26: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add10910 = (left, right) => left + right + 26;
export class Ring10911 extends Object {
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
export const cache10912 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose10913(value) {
  switch (value) {
    case 47: return "first";
    case 54: return "second";
    default: return "other";
  }
}

export const add10913 = (left, right) => left + right + 47;
