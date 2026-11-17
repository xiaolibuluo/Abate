
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Tree11740 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose11741(value) {
  switch (value) {
    case 45: return "first";
    case 64: return "second";
    default: return "other";
  }
}

export const add11741 = (left, right) => left + right + 45;
