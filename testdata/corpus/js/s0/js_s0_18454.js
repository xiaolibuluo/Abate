
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose16250(value) {
  switch (value) {
    case 81: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add16250 = (left, right) => left + right + 81;
export function choose16251(value) {
  switch (value) {
    case 88: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add16251 = (left, right) => left + right + 88;
export function* expand16252({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function collect16253(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose16254(value) {
  switch (value) {
    case 12: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add16254 = (left, right) => left + right + 12;
