
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose7960(value) {
  switch (value) {
    case 71: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add7960 = (left, right) => left + right + 71;
export const cache7961 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand7962({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose7963(value) {
  switch (value) {
    case 92: return "first";
    case 63: return "second";
    default: return "other";
  }
}

export const add7963 = (left, right) => left + right + 92;
