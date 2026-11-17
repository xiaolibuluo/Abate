
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Tree11440 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function collect11441(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose11442(value) {
  switch (value) {
    case 50: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add11442 = (left, right) => left + right + 50;
export function choose11443(value) {
  switch (value) {
    case 57: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add11443 = (left, right) => left + right + 57;
export const cache11444 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
