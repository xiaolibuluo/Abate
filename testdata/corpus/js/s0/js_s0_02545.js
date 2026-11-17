
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose2230(value) {
  switch (value) {
    case 91: return "first";
    case 45: return "second";
    default: return "other";
  }
}

export const add2230 = (left, right) => left + right + 91;
export function choose2231(value) {
  switch (value) {
    case 98: return "first";
    case 56: return "second";
    default: return "other";
  }
}

export const add2231 = (left, right) => left + right + 98;
export async function retry2232(task, attempts) {
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
export class Tree2233 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function* expand2234({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose2235(value) {
  switch (value) {
    case 29: return "first";
    case 39: return "second";
    default: return "other";
  }
}

export const add2235 = (left, right) => left + right + 29;
