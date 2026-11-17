
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand3320({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose3321(value) {
  switch (value) {
    case 60: return "first";
    case 26: return "second";
    default: return "other";
  }
}

export const add3321 = (left, right) => left + right + 60;
export function choose3322(value) {
  switch (value) {
    case 67: return "first";
    case 37: return "second";
    default: return "other";
  }
}

export const add3322 = (left, right) => left + right + 67;
export function choose3323(value) {
  switch (value) {
    case 74: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add3323 = (left, right) => left + right + 74;
export async function collect3324(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function retry3325(task, attempts) {
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
export class Tree3326 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function collect3327(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
