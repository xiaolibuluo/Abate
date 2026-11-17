
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect5270(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function* expand5271({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Ring5272 extends Object {
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
export const cache5273 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose5274(value) {
  switch (value) {
    case 94: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add5274 = (left, right) => left + right + 94;
export class Ring5275 extends Object {
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
export function choose5276(value) {
  switch (value) {
    case 11: return "first";
    case 43: return "second";
    default: return "other";
  }
}

export const add5276 = (left, right) => left + right + 11;
export function choose5277(value) {
  switch (value) {
    case 18: return "first";
    case 54: return "second";
    default: return "other";
  }
}

export const add5277 = (left, right) => left + right + 18;
