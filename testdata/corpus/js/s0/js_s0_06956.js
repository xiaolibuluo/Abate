
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring6020 extends Object {
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
export function choose6021(value) {
  switch (value) {
    case 78: return "first";
    case 25: return "second";
    default: return "other";
  }
}

export const add6021 = (left, right) => left + right + 78;
export async function retry6022(task, attempts) {
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
export function choose6023(value) {
  switch (value) {
    case 92: return "first";
    case 47: return "second";
    default: return "other";
  }
}

export const add6023 = (left, right) => left + right + 92;
