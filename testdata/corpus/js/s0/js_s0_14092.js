
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose12460(value) {
  switch (value) {
    case 101: return "first";
    case 69: return "second";
    default: return "other";
  }
}

export const add12460 = (left, right) => left + right + 101;
export class Tree12461 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function collect12462(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export const cache12463 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export const cache12464 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose12465(value) {
  switch (value) {
    case 39: return "first";
    case 63: return "second";
    default: return "other";
  }
}

export const add12465 = (left, right) => left + right + 39;
