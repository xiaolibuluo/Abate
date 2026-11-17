
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose3670(value) {
  switch (value) {
    case 23: return "first";
    case 16: return "second";
    default: return "other";
  }
}

export const add3670 = (left, right) => left + right + 23;
export async function retry3671(task, attempts) {
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
export function choose3672(value) {
  switch (value) {
    case 37: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add3672 = (left, right) => left + right + 37;
export class Tree3673 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function collect3674(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
