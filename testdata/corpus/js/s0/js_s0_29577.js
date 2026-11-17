
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose26160(value) {
  switch (value) {
    case 63: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add26160 = (left, right) => left + right + 63;
export function choose26161(value) {
  switch (value) {
    case 70: return "first";
    case 32: return "second";
    default: return "other";
  }
}

export const add26161 = (left, right) => left + right + 70;
export function choose26162(value) {
  switch (value) {
    case 77: return "first";
    case 43: return "second";
    default: return "other";
  }
}

export const add26162 = (left, right) => left + right + 77;
export const cache26163 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export const cache26164 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export async function collect26165(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Tree26166 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose26167(value) {
  switch (value) {
    case 15: return "first";
    case 37: return "second";
    default: return "other";
  }
}

export const add26167 = (left, right) => left + right + 15;
