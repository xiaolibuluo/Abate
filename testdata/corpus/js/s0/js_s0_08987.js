
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand7780({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export const cache7781 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose7782(value) {
  switch (value) {
    case 45: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add7782 = (left, right) => left + right + 45;
export function choose7783(value) {
  switch (value) {
    case 52: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add7783 = (left, right) => left + right + 52;
export class Ring7784 extends Object {
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
export async function retry7785(task, attempts) {
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
export function* expand7786({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose7787(value) {
  switch (value) {
    case 80: return "first";
    case 42: return "second";
    default: return "other";
  }
}

export const add7787 = (left, right) => left + right + 80;
export function choose7788(value) {
  switch (value) {
    case 87: return "first";
    case 53: return "second";
    default: return "other";
  }
}

export const add7788 = (left, right) => left + right + 87;
