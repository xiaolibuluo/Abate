
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose7230(value) {
  switch (value) {
    case 92: return "first";
    case 68: return "second";
    default: return "other";
  }
}

export const add7230 = (left, right) => left + right + 92;
export function choose7231(value) {
  switch (value) {
    case 99: return "first";
    case 18: return "second";
    default: return "other";
  }
}

export const add7231 = (left, right) => left + right + 99;
export function choose7232(value) {
  switch (value) {
    case 9: return "first";
    case 29: return "second";
    default: return "other";
  }
}

export const add7232 = (left, right) => left + right + 9;
export function* expand7233({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function collect7234(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Ring7235 extends Object {
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
export class Tree7236 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose7237(value) {
  switch (value) {
    case 44: return "first";
    case 23: return "second";
    default: return "other";
  }
}

export const add7237 = (left, right) => left + right + 44;
