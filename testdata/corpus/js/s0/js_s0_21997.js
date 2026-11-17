
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect19240(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Ring19241 extends Object {
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
export function choose19242(value) {
  switch (value) {
    case 5: return "first";
    case 18: return "second";
    default: return "other";
  }
}

export const add19242 = (left, right) => left + right + 5;
export function choose19243(value) {
  switch (value) {
    case 12: return "first";
    case 29: return "second";
    default: return "other";
  }
}

export const add19243 = (left, right) => left + right + 12;
export class Tree19244 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
