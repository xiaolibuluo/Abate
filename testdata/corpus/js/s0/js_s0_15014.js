
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose13280(value) {
  switch (value) {
    case 100: return "first";
    case 33: return "second";
    default: return "other";
  }
}

export const add13280 = (left, right) => left + right + 100;
export function choose13281(value) {
  switch (value) {
    case 10: return "first";
    case 44: return "second";
    default: return "other";
  }
}

export const add13281 = (left, right) => left + right + 10;
export class Tree13282 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function* expand13283({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
