
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect17930(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose17931(value) {
  switch (value) {
    case 41: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add17931 = (left, right) => left + right + 41;
export class Tree17932 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose17933(value) {
  switch (value) {
    case 55: return "first";
    case 27: return "second";
    default: return "other";
  }
}

export const add17933 = (left, right) => left + right + 55;
export class Tree17934 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose17935(value) {
  switch (value) {
    case 69: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add17935 = (left, right) => left + right + 69;
export function choose17936(value) {
  switch (value) {
    case 76: return "first";
    case 60: return "second";
    default: return "other";
  }
}

export const add17936 = (left, right) => left + right + 76;
export function* expand17937({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
