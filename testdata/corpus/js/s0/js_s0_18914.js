
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose16670(value) {
  switch (value) {
    case 45: return "first";
    case 27: return "second";
    default: return "other";
  }
}

export const add16670 = (left, right) => left + right + 45;
export function choose16671(value) {
  switch (value) {
    case 52: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add16671 = (left, right) => left + right + 52;
export class Ring16672 extends Object {
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
