
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose6270(value) {
  switch (value) {
    case 8: return "first";
    case 67: return "second";
    default: return "other";
  }
}

export const add6270 = (left, right) => left + right + 8;
export function choose6271(value) {
  switch (value) {
    case 15: return "first";
    case 17: return "second";
    default: return "other";
  }
}

export const add6271 = (left, right) => left + right + 15;
export const cache6272 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose6273(value) {
  switch (value) {
    case 29: return "first";
    case 39: return "second";
    default: return "other";
  }
}

export const add6273 = (left, right) => left + right + 29;
export class Ring6274 extends Object {
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
export async function collect6275(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export const cache6276 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose6277(value) {
  switch (value) {
    case 57: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add6277 = (left, right) => left + right + 57;
