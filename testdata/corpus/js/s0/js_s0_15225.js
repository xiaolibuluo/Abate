
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry13440(task, attempts) {
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
export function choose13441(value) {
  switch (value) {
    case 24: return "first";
    case 34: return "second";
    default: return "other";
  }
}

export const add13441 = (left, right) => left + right + 24;
export class Tree13442 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose13443(value) {
  switch (value) {
    case 38: return "first";
    case 56: return "second";
    default: return "other";
  }
}

export const add13443 = (left, right) => left + right + 38;
export function choose13444(value) {
  switch (value) {
    case 45: return "first";
    case 67: return "second";
    default: return "other";
  }
}

export const add13444 = (left, right) => left + right + 45;
export async function retry13445(task, attempts) {
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
