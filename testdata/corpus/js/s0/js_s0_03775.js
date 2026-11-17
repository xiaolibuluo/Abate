
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose3280(value) {
  switch (value) {
    case 98: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add3280 = (left, right) => left + right + 98;
export function choose3281(value) {
  switch (value) {
    case 8: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add3281 = (left, right) => left + right + 8;
export const cache3282 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose3283(value) {
  switch (value) {
    case 22: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add3283 = (left, right) => left + right + 22;
export const cache3284 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand3285({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export const cache3286 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export class Ring3287 extends Object {
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
