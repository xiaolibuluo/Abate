
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose1040(value) {
  switch (value) {
    case 96: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add1040 = (left, right) => left + right + 96;
export class Ring1041 extends Object {
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
