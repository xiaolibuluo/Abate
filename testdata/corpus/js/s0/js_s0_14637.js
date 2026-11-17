
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect12880(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose12881(value) {
  switch (value) {
    case 72: return "first";
    case 69: return "second";
    default: return "other";
  }
}

export const add12881 = (left, right) => left + right + 72;
export async function collect12882(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose12883(value) {
  switch (value) {
    case 86: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add12883 = (left, right) => left + right + 86;
export function* expand12884({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
