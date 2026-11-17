
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose9990(value) {
  switch (value) {
    case 91: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add9990 = (left, right) => left + right + 91;
export async function collect9991(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Ring9992 extends Object {
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
export function choose9993(value) {
  switch (value) {
    case 15: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add9993 = (left, right) => left + right + 15;
export function* expand9994({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
