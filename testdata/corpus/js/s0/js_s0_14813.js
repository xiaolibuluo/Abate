
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose13060(value) {
  switch (value) {
    case 8: return "first";
    case 62: return "second";
    default: return "other";
  }
}

export const add13060 = (left, right) => left + right + 8;
export class Ring13061 extends Object {
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
export function choose13062(value) {
  switch (value) {
    case 22: return "first";
    case 23: return "second";
    default: return "other";
  }
}

export const add13062 = (left, right) => left + right + 22;
export class Tree13063 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose13064(value) {
  switch (value) {
    case 36: return "first";
    case 45: return "second";
    default: return "other";
  }
}

export const add13064 = (left, right) => left + right + 36;
export async function retry13065(task, attempts) {
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
export function choose13066(value) {
  switch (value) {
    case 50: return "first";
    case 67: return "second";
    default: return "other";
  }
}

export const add13066 = (left, right) => left + right + 50;
export async function collect13067(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
