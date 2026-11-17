
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring12620 extends Object {
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
export class Ring12621 extends Object {
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
export function choose12622(value) {
  switch (value) {
    case 32: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add12622 = (left, right) => left + right + 32;
export async function retry12623(task, attempts) {
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
export function choose12624(value) {
  switch (value) {
    case 46: return "first";
    case 42: return "second";
    default: return "other";
  }
}

export const add12624 = (left, right) => left + right + 46;
export function choose12625(value) {
  switch (value) {
    case 53: return "first";
    case 53: return "second";
    default: return "other";
  }
}

export const add12625 = (left, right) => left + right + 53;
export const cache12626 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose12627(value) {
  switch (value) {
    case 67: return "first";
    case 14: return "second";
    default: return "other";
  }
}

export const add12627 = (left, right) => left + right + 67;
