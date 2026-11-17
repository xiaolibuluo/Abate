
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache15030 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export class Ring15031 extends Object {
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
export async function collect15032(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose15033(value) {
  switch (value) {
    case 68: return "first";
    case 10: return "second";
    default: return "other";
  }
}

export const add15033 = (left, right) => left + right + 68;
export function choose15034(value) {
  switch (value) {
    case 75: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add15034 = (left, right) => left + right + 75;
export class Ring15035 extends Object {
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
export function choose15036(value) {
  switch (value) {
    case 89: return "first";
    case 43: return "second";
    default: return "other";
  }
}

export const add15036 = (left, right) => left + right + 89;
export function* expand15037({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
