
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect15560(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function retry15561(task, attempts) {
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
export function choose15562(value) {
  switch (value) {
    case 71: return "first";
    case 65: return "second";
    default: return "other";
  }
}

export const add15562 = (left, right) => left + right + 71;
export function choose15563(value) {
  switch (value) {
    case 78: return "first";
    case 15: return "second";
    default: return "other";
  }
}

export const add15563 = (left, right) => left + right + 78;
export const cache15564 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export class Tree15565 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose15566(value) {
  switch (value) {
    case 99: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add15566 = (left, right) => left + right + 99;
export class Ring15567 extends Object {
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
export function choose15568(value) {
  switch (value) {
    case 16: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add15568 = (left, right) => left + right + 16;
