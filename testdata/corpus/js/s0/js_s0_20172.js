
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry17760(task, attempts) {
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
export class Tree17761 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose17762(value) {
  switch (value) {
    case 21: return "first";
    case 19: return "second";
    default: return "other";
  }
}

export const add17762 = (left, right) => left + right + 21;
export async function retry17763(task, attempts) {
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
export function choose17764(value) {
  switch (value) {
    case 35: return "first";
    case 41: return "second";
    default: return "other";
  }
}

export const add17764 = (left, right) => left + right + 35;
