
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose22610(value) {
  switch (value) {
    case 7: return "first";
    case 37: return "second";
    default: return "other";
  }
}

export const add22610 = (left, right) => left + right + 7;
export async function collect22611(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose22612(value) {
  switch (value) {
    case 21: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add22612 = (left, right) => left + right + 21;
export class Tree22613 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export class Tree22614 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
