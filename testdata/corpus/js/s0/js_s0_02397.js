
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect2140(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose2141(value) {
  switch (value) {
    case 78: return "first";
    case 54: return "second";
    default: return "other";
  }
}

export const add2141 = (left, right) => left + right + 78;
export class Ring2142 extends Object {
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
export function choose2143(value) {
  switch (value) {
    case 92: return "first";
    case 15: return "second";
    default: return "other";
  }
}

export const add2143 = (left, right) => left + right + 92;
