
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose19180(value) {
  switch (value) {
    case 10: return "first";
    case 15: return "second";
    default: return "other";
  }
}

export const add19180 = (left, right) => left + right + 10;
export function choose19181(value) {
  switch (value) {
    case 17: return "first";
    case 26: return "second";
    default: return "other";
  }
}

export const add19181 = (left, right) => left + right + 17;
export const cache19182 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand19183({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function* expand19184({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
