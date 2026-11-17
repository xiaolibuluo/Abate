
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose10170(value) {
  switch (value) {
    case 34: return "first";
    case 52: return "second";
    default: return "other";
  }
}

export const add10170 = (left, right) => left + right + 34;
export async function collect10171(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function collect10172(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function retry10173(task, attempts) {
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
export async function retry10174(task, attempts) {
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
export function choose10175(value) {
  switch (value) {
    case 69: return "first";
    case 46: return "second";
    default: return "other";
  }
}

export const add10175 = (left, right) => left + right + 69;
export class Ring10176 extends Object {
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
export function choose10177(value) {
  switch (value) {
    case 83: return "first";
    case 68: return "second";
    default: return "other";
  }
}

export const add10177 = (left, right) => left + right + 83;
