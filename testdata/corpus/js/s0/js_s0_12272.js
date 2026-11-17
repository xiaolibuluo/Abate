
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose11090(value) {
  switch (value) {
    case 66: return "first";
    case 25: return "second";
    default: return "other";
  }
}

export const add11090 = (left, right) => left + right + 66;
export function choose11091(value) {
  switch (value) {
    case 73: return "first";
    case 36: return "second";
    default: return "other";
  }
}

export const add11091 = (left, right) => left + right + 73;
export async function retry11092(task, attempts) {
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
export function choose11093(value) {
  switch (value) {
    case 87: return "first";
    case 58: return "second";
    default: return "other";
  }
}

export const add11093 = (left, right) => left + right + 87;
export async function collect11094(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function collect11095(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function* expand11096({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Ring11097 extends Object {
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
