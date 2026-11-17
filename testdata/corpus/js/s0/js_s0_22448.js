
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring19700 extends Object {
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
export function choose19701(value) {
  switch (value) {
    case 14: return "first";
    case 24: return "second";
    default: return "other";
  }
}

export const add19701 = (left, right) => left + right + 14;
export async function collect19702(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose19703(value) {
  switch (value) {
    case 28: return "first";
    case 46: return "second";
    default: return "other";
  }
}

export const add19703 = (left, right) => left + right + 28;
export function* expand19704({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
