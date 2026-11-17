
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Tree10250 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose10251(value) {
  switch (value) {
    case 48: return "first";
    case 58: return "second";
    default: return "other";
  }
}

export const add10251 = (left, right) => left + right + 48;
