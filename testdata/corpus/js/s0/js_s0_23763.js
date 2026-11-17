
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose20810(value) {
  switch (value) {
    case 92: return "first";
    case 58: return "second";
    default: return "other";
  }
}

export const add20810 = (left, right) => left + right + 92;
export function choose20811(value) {
  switch (value) {
    case 99: return "first";
    case 69: return "second";
    default: return "other";
  }
}

export const add20811 = (left, right) => left + right + 99;
export function* expand20812({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose20813(value) {
  switch (value) {
    case 16: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add20813 = (left, right) => left + right + 16;
export function* expand20814({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function* expand20815({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Ring20816 extends Object {
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
export class Tree20817 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
