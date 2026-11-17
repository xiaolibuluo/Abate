
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand4530({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose4531(value) {
  switch (value) {
    case 81: return "first";
    case 19: return "second";
    default: return "other";
  }
}

export const add4531 = (left, right) => left + right + 81;
export function choose4532(value) {
  switch (value) {
    case 88: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add4532 = (left, right) => left + right + 88;
export const cache4533 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
