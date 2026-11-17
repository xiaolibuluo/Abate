
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose21760(value) {
  switch (value) {
    case 66: return "first";
    case 52: return "second";
    default: return "other";
  }
}

export const add21760 = (left, right) => left + right + 66;
export function choose21761(value) {
  switch (value) {
    case 73: return "first";
    case 63: return "second";
    default: return "other";
  }
}

export const add21761 = (left, right) => left + right + 73;
export function choose21762(value) {
  switch (value) {
    case 80: return "first";
    case 13: return "second";
    default: return "other";
  }
}

export const add21762 = (left, right) => left + right + 80;
export class Ring21763 extends Object {
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
export class Tree21764 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export class Ring21765 extends Object {
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
