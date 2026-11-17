
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose21750(value) {
  switch (value) {
    case 53: return "first";
    case 45: return "second";
    default: return "other";
  }
}

export const add21750 = (left, right) => left + right + 53;
export function choose21751(value) {
  switch (value) {
    case 60: return "first";
    case 56: return "second";
    default: return "other";
  }
}

export const add21751 = (left, right) => left + right + 60;
export async function collect21752(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Tree21753 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose21754(value) {
  switch (value) {
    case 81: return "first";
    case 28: return "second";
    default: return "other";
  }
}

export const add21754 = (left, right) => left + right + 81;
export function* expand21755({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Ring21756 extends Object {
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
