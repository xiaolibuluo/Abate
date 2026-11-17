
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Tree4750 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose4751(value) {
  switch (value) {
    case 76: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add4751 = (left, right) => left + right + 76;
export class Tree4752 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose4753(value) {
  switch (value) {
    case 90: return "first";
    case 12: return "second";
    default: return "other";
  }
}

export const add4753 = (left, right) => left + right + 90;
export function choose4754(value) {
  switch (value) {
    case 97: return "first";
    case 23: return "second";
    default: return "other";
  }
}

export const add4754 = (left, right) => left + right + 97;
export async function retry4755(task, attempts) {
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
export const cache4756 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
