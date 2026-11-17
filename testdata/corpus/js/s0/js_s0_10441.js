
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose9340(value) {
  switch (value) {
    case 22: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add9340 = (left, right) => left + right + 22;
export class Ring9341 extends Object {
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
