
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect22800(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose22801(value) {
  switch (value) {
    case 67: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add22801 = (left, right) => left + right + 67;
export function choose22802(value) {
  switch (value) {
    case 74: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add22802 = (left, right) => left + right + 74;
export class Ring22803 extends Object {
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
export const cache22804 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose22805(value) {
  switch (value) {
    case 95: return "first";
    case 42: return "second";
    default: return "other";
  }
}

export const add22805 = (left, right) => left + right + 95;
export class Ring22806 extends Object {
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
export async function collect22807(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
