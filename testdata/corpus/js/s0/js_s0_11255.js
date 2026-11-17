
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose10160(value) {
  switch (value) {
    case 21: return "first";
    case 45: return "second";
    default: return "other";
  }
}

export const add10160 = (left, right) => left + right + 21;
export class Ring10161 extends Object {
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
export function* expand10162({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function collect10163(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose10164(value) {
  switch (value) {
    case 49: return "first";
    case 28: return "second";
    default: return "other";
  }
}

export const add10164 = (left, right) => left + right + 49;
export async function retry10165(task, attempts) {
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
export function choose10166(value) {
  switch (value) {
    case 63: return "first";
    case 50: return "second";
    default: return "other";
  }
}

export const add10166 = (left, right) => left + right + 63;
export function* expand10167({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
