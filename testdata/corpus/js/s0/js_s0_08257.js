
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect7270(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Tree7271 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export class Tree7272 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function* expand7273({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose7274(value) {
  switch (value) {
    case 75: return "first";
    case 18: return "second";
    default: return "other";
  }
}

export const add7274 = (left, right) => left + right + 75;
export function choose7275(value) {
  switch (value) {
    case 82: return "first";
    case 29: return "second";
    default: return "other";
  }
}

export const add7275 = (left, right) => left + right + 82;
export class Ring7276 extends Object {
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
export function choose7277(value) {
  switch (value) {
    case 96: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add7277 = (left, right) => left + right + 96;
export function choose7278(value) {
  switch (value) {
    case 6: return "first";
    case 62: return "second";
    default: return "other";
  }
}

export const add7278 = (left, right) => left + right + 6;
