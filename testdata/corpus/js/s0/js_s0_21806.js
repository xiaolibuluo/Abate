
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose19030(value) {
  switch (value) {
    case 9: return "first";
    case 32: return "second";
    default: return "other";
  }
}

export const add19030 = (left, right) => left + right + 9;
export class Tree19031 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
