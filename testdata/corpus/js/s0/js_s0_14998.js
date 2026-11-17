
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose13240(value) {
  switch (value) {
    case 48: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add13240 = (left, right) => left + right + 48;
export class Ring13241 extends Object {
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
export function choose13242(value) {
  switch (value) {
    case 62: return "first";
    case 27: return "second";
    default: return "other";
  }
}

export const add13242 = (left, right) => left + right + 62;
