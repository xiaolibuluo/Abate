
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Ring22660 extends Object {
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
export function choose22661(value) {
  switch (value) {
    case 79: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add22661 = (left, right) => left + right + 79;
export function choose22662(value) {
  switch (value) {
    case 86: return "first";
    case 33: return "second";
    default: return "other";
  }
}

export const add22662 = (left, right) => left + right + 86;
export class Tree22663 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function collect22664(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
