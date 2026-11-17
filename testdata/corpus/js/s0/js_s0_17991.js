
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring15830 extends Object {
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
export function choose15831(value) {
  switch (value) {
    case 27: return "first";
    case 60: return "second";
    default: return "other";
  }
}

export const add15831 = (left, right) => left + right + 27;
export function choose15832(value) {
  switch (value) {
    case 34: return "first";
    case 10: return "second";
    default: return "other";
  }
}

export const add15832 = (left, right) => left + right + 34;
export async function collect15833(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose15834(value) {
  switch (value) {
    case 48: return "first";
    case 32: return "second";
    default: return "other";
  }
}

export const add15834 = (left, right) => left + right + 48;
