
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose14840(value) {
  switch (value) {
    case 91: return "first";
    case 27: return "second";
    default: return "other";
  }
}

export const add14840 = (left, right) => left + right + 91;
export function choose14841(value) {
  switch (value) {
    case 98: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add14841 = (left, right) => left + right + 98;
export function choose14842(value) {
  switch (value) {
    case 8: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add14842 = (left, right) => left + right + 8;
export function choose14843(value) {
  switch (value) {
    case 15: return "first";
    case 60: return "second";
    default: return "other";
  }
}

export const add14843 = (left, right) => left + right + 15;
export function* expand14844({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function collect14845(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
