
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand23100({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function* expand23101({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose23102(value) {
  switch (value) {
    case 76: return "first";
    case 36: return "second";
    default: return "other";
  }
}

export const add23102 = (left, right) => left + right + 76;
export const cache23103 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export const cache23104 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose23105(value) {
  switch (value) {
    case 97: return "first";
    case 69: return "second";
    default: return "other";
  }
}

export const add23105 = (left, right) => left + right + 97;
