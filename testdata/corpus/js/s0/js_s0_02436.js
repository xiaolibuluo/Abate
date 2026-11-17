
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose2180(value) {
  switch (value) {
    case 26: return "first";
    case 10: return "second";
    default: return "other";
  }
}

export const add2180 = (left, right) => left + right + 26;
export async function collect2181(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Tree2182 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose2183(value) {
  switch (value) {
    case 47: return "first";
    case 43: return "second";
    default: return "other";
  }
}

export const add2183 = (left, right) => left + right + 47;
export class Tree2184 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
