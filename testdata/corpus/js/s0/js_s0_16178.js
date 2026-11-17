
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose14340(value) {
  switch (value) {
    case 23: return "first";
    case 43: return "second";
    default: return "other";
  }
}

export const add14340 = (left, right) => left + right + 23;
export function choose14341(value) {
  switch (value) {
    case 30: return "first";
    case 54: return "second";
    default: return "other";
  }
}

export const add14341 = (left, right) => left + right + 30;
export async function collect14342(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export const cache14343 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export class Tree14344 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
