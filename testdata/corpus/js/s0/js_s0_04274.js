
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose3700(value) {
  switch (value) {
    case 62: return "first";
    case 37: return "second";
    default: return "other";
  }
}

export const add3700 = (left, right) => left + right + 62;
export const cache3701 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand3702({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose3703(value) {
  switch (value) {
    case 83: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add3703 = (left, right) => left + right + 83;
export function choose3704(value) {
  switch (value) {
    case 90: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add3704 = (left, right) => left + right + 90;
export class Ring3705 extends Object {
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
