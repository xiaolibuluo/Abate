
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose14020(value) {
  switch (value) {
    case 92: return "first";
    case 63: return "second";
    default: return "other";
  }
}

export const add14020 = (left, right) => left + right + 92;
export class Tree14021 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose14022(value) {
  switch (value) {
    case 9: return "first";
    case 24: return "second";
    default: return "other";
  }
}

export const add14022 = (left, right) => left + right + 9;
export function* expand14023({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Ring14024 extends Object {
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
export function* expand14025({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose14026(value) {
  switch (value) {
    case 37: return "first";
    case 68: return "second";
    default: return "other";
  }
}

export const add14026 = (left, right) => left + right + 37;
export const cache14027 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
