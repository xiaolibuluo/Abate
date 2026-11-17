
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose22560(value) {
  switch (value) {
    case 39: return "first";
    case 63: return "second";
    default: return "other";
  }
}

export const add22560 = (left, right) => left + right + 39;
export function choose22561(value) {
  switch (value) {
    case 46: return "first";
    case 13: return "second";
    default: return "other";
  }
}

export const add22561 = (left, right) => left + right + 46;
export function* expand22562({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Tree22563 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
