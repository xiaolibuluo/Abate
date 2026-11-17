
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose3100(value) {
  switch (value) {
    case 58: return "first";
    case 44: return "second";
    default: return "other";
  }
}

export const add3100 = (left, right) => left + right + 58;
export async function collect3101(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose3102(value) {
  switch (value) {
    case 72: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add3102 = (left, right) => left + right + 72;
export class Tree3103 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
