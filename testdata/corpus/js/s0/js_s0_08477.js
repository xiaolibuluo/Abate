
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose7410(value) {
  switch (value) {
    case 35: return "first";
    case 11: return "second";
    default: return "other";
  }
}

export const add7410 = (left, right) => left + right + 35;
export function choose7411(value) {
  switch (value) {
    case 42: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add7411 = (left, right) => left + right + 42;
export class Tree7412 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
