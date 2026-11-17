
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry12920(task, attempts) {
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
export function choose12921(value) {
  switch (value) {
    case 27: return "first";
    case 36: return "second";
    default: return "other";
  }
}

export const add12921 = (left, right) => left + right + 27;
export function choose12922(value) {
  switch (value) {
    case 34: return "first";
    case 47: return "second";
    default: return "other";
  }
}

export const add12922 = (left, right) => left + right + 34;
export function choose12923(value) {
  switch (value) {
    case 41: return "first";
    case 58: return "second";
    default: return "other";
  }
}

export const add12923 = (left, right) => left + right + 41;
export class Tree12924 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function collect12925(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
