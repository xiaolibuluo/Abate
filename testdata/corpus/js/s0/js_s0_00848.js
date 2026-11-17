
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand790({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose791(value) {
  switch (value) {
    case 69: return "first";
    case 24: return "second";
    default: return "other";
  }
}

export const add791 = (left, right) => left + right + 69;
export class Ring792 extends Object {
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
export function choose793(value) {
  switch (value) {
    case 83: return "first";
    case 46: return "second";
    default: return "other";
  }
}

export const add793 = (left, right) => left + right + 83;
export function choose794(value) {
  switch (value) {
    case 90: return "first";
    case 57: return "second";
    default: return "other";
  }
}

export const add794 = (left, right) => left + right + 90;
export function* expand795({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
