
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose18730(value) {
  switch (value) {
    case 7: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add18730 = (left, right) => left + right + 7;
export async function collect18731(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function* expand18732({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function* expand18733({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose18734(value) {
  switch (value) {
    case 35: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add18734 = (left, right) => left + right + 35;
export async function collect18735(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
