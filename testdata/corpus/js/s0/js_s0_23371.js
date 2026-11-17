
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose20520(value) {
  switch (value) {
    case 6: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add20520 = (left, right) => left + right + 6;
export function choose20521(value) {
  switch (value) {
    case 13: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add20521 = (left, right) => left + right + 13;
export class Ring20522 extends Object {
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
