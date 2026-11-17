
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose24120(value) {
  switch (value) {
    case 30: return "first";
    case 57: return "second";
    default: return "other";
  }
}

export const add24120 = (left, right) => left + right + 30;
export function choose24121(value) {
  switch (value) {
    case 37: return "first";
    case 68: return "second";
    default: return "other";
  }
}

export const add24121 = (left, right) => left + right + 37;
export async function collect24122(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export class Tree24123 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function* expand24124({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose24125(value) {
  switch (value) {
    case 65: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add24125 = (left, right) => left + right + 65;
