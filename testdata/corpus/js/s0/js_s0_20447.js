
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose17970(value) {
  switch (value) {
    case 86: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add17970 = (left, right) => left + right + 86;
export class Ring17971 extends Object {
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
export class Tree17972 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose17973(value) {
  switch (value) {
    case 10: return "first";
    case 55: return "second";
    default: return "other";
  }
}

export const add17973 = (left, right) => left + right + 10;
export async function retry17974(task, attempts) {
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
export function choose17975(value) {
  switch (value) {
    case 24: return "first";
    case 16: return "second";
    default: return "other";
  }
}

export const add17975 = (left, right) => left + right + 24;
export const cache17976 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose17977(value) {
  switch (value) {
    case 38: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add17977 = (left, right) => left + right + 38;
