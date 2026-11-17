
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose7710(value) {
  switch (value) {
    case 37: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add7710 = (left, right) => left + right + 37;
export function choose7711(value) {
  switch (value) {
    case 44: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add7711 = (left, right) => left + right + 44;
export function choose7712(value) {
  switch (value) {
    case 51: return "first";
    case 60: return "second";
    default: return "other";
  }
}

export const add7712 = (left, right) => left + right + 51;
export class Tree7713 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function retry7714(task, attempts) {
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
export class Ring7715 extends Object {
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
export async function retry7716(task, attempts) {
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
