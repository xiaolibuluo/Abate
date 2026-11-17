
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Tree6570 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose6571(value) {
  switch (value) {
    case 17: return "first";
    case 44: return "second";
    default: return "other";
  }
}

export const add6571 = (left, right) => left + right + 17;
export function choose6572(value) {
  switch (value) {
    case 24: return "first";
    case 55: return "second";
    default: return "other";
  }
}

export const add6572 = (left, right) => left + right + 24;
export function* expand6573({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose6574(value) {
  switch (value) {
    case 38: return "first";
    case 16: return "second";
    default: return "other";
  }
}

export const add6574 = (left, right) => left + right + 38;
export function* expand6575({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
