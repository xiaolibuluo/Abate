
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose10180(value) {
  switch (value) {
    case 47: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add10180 = (left, right) => left + right + 47;
export function choose10181(value) {
  switch (value) {
    case 54: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add10181 = (left, right) => left + right + 54;
export async function collect10182(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose10183(value) {
  switch (value) {
    case 68: return "first";
    case 31: return "second";
    default: return "other";
  }
}

export const add10183 = (left, right) => left + right + 68;
export async function retry10184(task, attempts) {
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
export function choose10185(value) {
  switch (value) {
    case 82: return "first";
    case 53: return "second";
    default: return "other";
  }
}

export const add10185 = (left, right) => left + right + 82;
export class Ring10186 extends Object {
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
export class Ring10187 extends Object {
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
