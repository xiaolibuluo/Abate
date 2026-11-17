
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring20620 extends Object {
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
export const cache20621 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose20622(value) {
  switch (value) {
    case 53: return "first";
    case 69: return "second";
    default: return "other";
  }
}

export const add20622 = (left, right) => left + right + 53;
export function choose20623(value) {
  switch (value) {
    case 60: return "first";
    case 19: return "second";
    default: return "other";
  }
}

export const add20623 = (left, right) => left + right + 60;
