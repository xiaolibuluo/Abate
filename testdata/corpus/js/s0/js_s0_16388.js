
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Tree14500 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export const cache14501 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose14502(value) {
  switch (value) {
    case 51: return "first";
    case 55: return "second";
    default: return "other";
  }
}

export const add14502 = (left, right) => left + right + 51;
export async function retry14503(task, attempts) {
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
export function* expand14504({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Tree14505 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose14506(value) {
  switch (value) {
    case 79: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add14506 = (left, right) => left + right + 79;
export function choose14507(value) {
  switch (value) {
    case 86: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add14507 = (left, right) => left + right + 86;
