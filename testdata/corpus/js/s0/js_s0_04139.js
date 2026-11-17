
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose3620(value) {
  switch (value) {
    case 55: return "first";
    case 42: return "second";
    default: return "other";
  }
}

export const add3620 = (left, right) => left + right + 55;
export async function collect3621(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose3622(value) {
  switch (value) {
    case 69: return "first";
    case 64: return "second";
    default: return "other";
  }
}

export const add3622 = (left, right) => left + right + 69;
export class Tree3623 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export const cache3624 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
