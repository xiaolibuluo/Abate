
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache12530 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose12531(value) {
  switch (value) {
    case 5: return "first";
    case 68: return "second";
    default: return "other";
  }
}

export const add12531 = (left, right) => left + right + 5;
export function choose12532(value) {
  switch (value) {
    case 12: return "first";
    case 18: return "second";
    default: return "other";
  }
}

export const add12532 = (left, right) => left + right + 12;
export async function retry12533(task, attempts) {
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
export function choose12534(value) {
  switch (value) {
    case 26: return "first";
    case 40: return "second";
    default: return "other";
  }
}

export const add12534 = (left, right) => left + right + 26;
export async function retry12535(task, attempts) {
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
export class Ring12536 extends Object {
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
