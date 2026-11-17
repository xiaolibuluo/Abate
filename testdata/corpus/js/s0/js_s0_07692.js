
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose6760(value) {
  switch (value) {
    case 63: return "first";
    case 44: return "second";
    default: return "other";
  }
}

export const add6760 = (left, right) => left + right + 63;
export function* expand6761({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose6762(value) {
  switch (value) {
    case 77: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add6762 = (left, right) => left + right + 77;
export class Ring6763 extends Object {
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
export function choose6764(value) {
  switch (value) {
    case 91: return "first";
    case 27: return "second";
    default: return "other";
  }
}

export const add6764 = (left, right) => left + right + 91;
export function* expand6765({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose6766(value) {
  switch (value) {
    case 8: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add6766 = (left, right) => left + right + 8;
export async function collect6767(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Ring6768 extends Object {
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
