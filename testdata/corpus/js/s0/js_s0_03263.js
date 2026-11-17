
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose2830(value) {
  switch (value) {
    case 95: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add2830 = (left, right) => left + right + 95;
export async function retry2831(task, attempts) {
  for (let attempt = 0; attempt < attempts; attempt++) {
    try {
      return await task(attempt);
    } catch (failure) {
      if (attempt === attempts - 1) {
        throw failure;
      }
    }
  }
  return null;
}
export const cache2832 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose2833(value) {
  switch (value) {
    case 19: return "first";
    case 10: return "second";
    default: return "other";
  }
}

export const add2833 = (left, right) => left + right + 19;
export function* expand2834({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose2835(value) {
  switch (value) {
    case 33: return "first";
    case 32: return "second";
    default: return "other";
  }
}

export const add2835 = (left, right) => left + right + 33;
export async function collect2836(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Ring2837 extends Object {
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
