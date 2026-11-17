
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose14830(value) {
  switch (value) {
    case 78: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add14830 = (left, right) => left + right + 78;
export async function collect14831(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose14832(value) {
  switch (value) {
    case 92: return "first";
    case 42: return "second";
    default: return "other";
  }
}

export const add14832 = (left, right) => left + right + 92;
export function choose14833(value) {
  switch (value) {
    case 99: return "first";
    case 53: return "second";
    default: return "other";
  }
}

export const add14833 = (left, right) => left + right + 99;
export function* expand14834({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
