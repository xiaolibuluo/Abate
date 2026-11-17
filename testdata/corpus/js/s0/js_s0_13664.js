
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand12110({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose12111(value) {
  switch (value) {
    case 41: return "first";
    case 18: return "second";
    default: return "other";
  }
}

export const add12111 = (left, right) => left + right + 41;
export function choose12112(value) {
  switch (value) {
    case 48: return "first";
    case 29: return "second";
    default: return "other";
  }
}

export const add12112 = (left, right) => left + right + 48;
export async function collect12113(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function* expand12114({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export class Tree12115 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export class Tree12116 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export class Tree12117 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose12118(value) {
  switch (value) {
    case 90: return "first";
    case 34: return "second";
    default: return "other";
  }
}

export const add12118 = (left, right) => left + right + 90;
