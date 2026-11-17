
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring25050 extends Object {
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
export async function collect25051(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose25052(value) {
  switch (value) {
    case 89: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add25052 = (left, right) => left + right + 89;
export function choose25053(value) {
  switch (value) {
    case 96: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add25053 = (left, right) => left + right + 96;
export function choose25054(value) {
  switch (value) {
    case 6: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add25054 = (left, right) => left + right + 6;
