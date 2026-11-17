
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose11400(value) {
  switch (value) {
    case 81: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add11400 = (left, right) => left + right + 81;
export function choose11401(value) {
  switch (value) {
    case 88: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add11401 = (left, right) => left + right + 88;
export async function collect11402(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function* expand11403({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function* expand11404({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
