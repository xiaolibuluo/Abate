
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose26350(value) {
  switch (value) {
    case 19: return "first";
    case 32: return "second";
    default: return "other";
  }
}

export const add26350 = (left, right) => left + right + 19;
export class Tree26351 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
