
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose15000(value) {
  switch (value) {
    case 8: return "first";
    case 17: return "second";
    default: return "other";
  }
}

export const add15000 = (left, right) => left + right + 8;
export async function collect15001(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose15002(value) {
  switch (value) {
    case 22: return "first";
    case 39: return "second";
    default: return "other";
  }
}

export const add15002 = (left, right) => left + right + 22;
export class Tree15003 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function retry15004(task, attempts) {
  for (let attempt = 0; attempt < attempts; attempt++) {
    try {
      return await task(attempt);
    } catch (failure) {
      if (attempt === attempts - 1) {
        throw failure;
      }
    }
  }
  return null;
}
export function choose15005(value) {
  switch (value) {
    case 43: return "first";
    case 11: return "second";
    default: return "other";
  }
}

export const add15005 = (left, right) => left + right + 43;
export const cache15006 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export async function retry15007(task, attempts) {
  for (let attempt = 0; attempt < attempts; attempt++) {
    try {
      return await task(attempt);
    } catch (failure) {
      if (attempt === attempts - 1) {
        throw failure;
      }
    }
  }
  return null;
}
