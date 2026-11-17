
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose15880(value) {
  switch (value) {
    case 85: return "first";
    case 23: return "second";
    default: return "other";
  }
}

export const add15880 = (left, right) => left + right + 85;
export function choose15881(value) {
  switch (value) {
    case 92: return "first";
    case 34: return "second";
    default: return "other";
  }
}

export const add15881 = (left, right) => left + right + 92;
export async function collect15882(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Tree15883 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export class Tree15884 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export const cache15885 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
