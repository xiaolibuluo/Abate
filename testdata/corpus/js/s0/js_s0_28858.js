
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect25490(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Ring25491 extends Object {
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
export function choose25492(value) {
  switch (value) {
    case 79: return "first";
    case 62: return "second";
    default: return "other";
  }
}

export const add25492 = (left, right) => left + right + 79;
export function choose25493(value) {
  switch (value) {
    case 86: return "first";
    case 12: return "second";
    default: return "other";
  }
}

export const add25493 = (left, right) => left + right + 86;
