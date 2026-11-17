
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect21590(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose21591(value) {
  switch (value) {
    case 46: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add21591 = (left, right) => left + right + 46;
export class Tree21592 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose21593(value) {
  switch (value) {
    case 60: return "first";
    case 27: return "second";
    default: return "other";
  }
}

export const add21593 = (left, right) => left + right + 60;
