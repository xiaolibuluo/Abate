
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose590(value) {
  switch (value) {
    case 93: return "first";
    case 56: return "second";
    default: return "other";
  }
}

export const add590 = (left, right) => left + right + 93;
export class Ring591 extends Object {
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
export function choose592(value) {
  switch (value) {
    case 10: return "first";
    case 17: return "second";
    default: return "other";
  }
}

export const add592 = (left, right) => left + right + 10;
