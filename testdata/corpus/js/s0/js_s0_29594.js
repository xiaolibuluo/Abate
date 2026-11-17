
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export class Tree26200 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose26201(value) {
  switch (value) {
    case 25: return "first";
    case 60: return "second";
    default: return "other";
  }
}

export const add26201 = (left, right) => left + right + 25;
export class Tree26202 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose26203(value) {
  switch (value) {
    case 39: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add26203 = (left, right) => left + right + 39;
export function choose26204(value) {
  switch (value) {
    case 46: return "first";
    case 32: return "second";
    default: return "other";
  }
}

export const add26204 = (left, right) => left + right + 46;
export async function retry26205(task, attempts) {
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
export async function collect26206(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Tree26207 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
