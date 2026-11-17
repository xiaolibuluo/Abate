
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose26080(value) {
  switch (value) {
    case 56: return "first";
    case 26: return "second";
    default: return "other";
  }
}

export const add26080 = (left, right) => left + right + 56;
export class Tree26081 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose26082(value) {
  switch (value) {
    case 70: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add26082 = (left, right) => left + right + 70;
export class Tree26083 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function collect26084(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export const cache26085 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
