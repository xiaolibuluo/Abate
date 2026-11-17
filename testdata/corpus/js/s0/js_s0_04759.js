
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring4140 extends Object {
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
export function choose4141(value) {
  switch (value) {
    case 59: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add4141 = (left, right) => left + right + 59;
export function* expand4142({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose4143(value) {
  switch (value) {
    case 73: return "first";
    case 12: return "second";
    default: return "other";
  }
}

export const add4143 = (left, right) => left + right + 73;
export async function retry4144(task, attempts) {
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
