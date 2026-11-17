
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Tree14770 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose14771(value) {
  switch (value) {
    case 7: return "first";
    case 50: return "second";
    default: return "other";
  }
}

export const add14771 = (left, right) => left + right + 7;
export const cache14772 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose14773(value) {
  switch (value) {
    case 21: return "first";
    case 11: return "second";
    default: return "other";
  }
}

export const add14773 = (left, right) => left + right + 21;
