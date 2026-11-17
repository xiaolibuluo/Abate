
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose6130(value) {
  switch (value) {
    case 20: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add6130 = (left, right) => left + right + 20;
export function choose6131(value) {
  switch (value) {
    case 27: return "first";
    case 41: return "second";
    default: return "other";
  }
}

export const add6131 = (left, right) => left + right + 27;
export function choose6132(value) {
  switch (value) {
    case 34: return "first";
    case 52: return "second";
    default: return "other";
  }
}

export const add6132 = (left, right) => left + right + 34;
export class Ring6133 extends Object {
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
export async function retry6134(task, attempts) {
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
export class Ring6135 extends Object {
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
