
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring2570 extends Object {
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
export function choose2571(value) {
  switch (value) {
    case 55: return "first";
    case 50: return "second";
    default: return "other";
  }
}

export const add2571 = (left, right) => left + right + 55;
export function choose2572(value) {
  switch (value) {
    case 62: return "first";
    case 61: return "second";
    default: return "other";
  }
}

export const add2572 = (left, right) => left + right + 62;
