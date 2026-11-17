
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose25720(value) {
  switch (value) {
    case 73: return "first";
    case 18: return "second";
    default: return "other";
  }
}

export const add25720 = (left, right) => left + right + 73;
export class Ring25721 extends Object {
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
export function choose25722(value) {
  switch (value) {
    case 87: return "first";
    case 40: return "second";
    default: return "other";
  }
}

export const add25722 = (left, right) => left + right + 87;
export function choose25723(value) {
  switch (value) {
    case 94: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add25723 = (left, right) => left + right + 94;
export function* expand25724({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export const cache25725 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose25726(value) {
  switch (value) {
    case 18: return "first";
    case 23: return "second";
    default: return "other";
  }
}

export const add25726 = (left, right) => left + right + 18;
export class Ring25727 extends Object {
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
