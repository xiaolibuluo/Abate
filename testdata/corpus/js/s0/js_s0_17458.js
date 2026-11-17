
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose15400(value) {
  switch (value) {
    case 43: return "first";
    case 53: return "second";
    default: return "other";
  }
}

export const add15400 = (left, right) => left + right + 43;
export function* expand15401({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose15402(value) {
  switch (value) {
    case 57: return "first";
    case 14: return "second";
    default: return "other";
  }
}

export const add15402 = (left, right) => left + right + 57;
export async function retry15403(task, attempts) {
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
export function choose15404(value) {
  switch (value) {
    case 71: return "first";
    case 36: return "second";
    default: return "other";
  }
}

export const add15404 = (left, right) => left + right + 71;
export class Ring15405 extends Object {
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
