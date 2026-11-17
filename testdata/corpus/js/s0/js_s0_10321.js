
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring9170 extends Object {
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
export async function retry9171(task, attempts) {
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
export async function collect9172(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose9173(value) {
  switch (value) {
    case 16: return "first";
    case 56: return "second";
    default: return "other";
  }
}

export const add9173 = (left, right) => left + right + 16;
export class Ring9174 extends Object {
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
export function choose9175(value) {
  switch (value) {
    case 30: return "first";
    case 17: return "second";
    default: return "other";
  }
}

export const add9175 = (left, right) => left + right + 30;
export function choose9176(value) {
  switch (value) {
    case 37: return "first";
    case 28: return "second";
    default: return "other";
  }
}

export const add9176 = (left, right) => left + right + 37;
