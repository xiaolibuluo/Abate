
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand16180({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export const cache16181 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose16182(value) {
  switch (value) {
    case 101: return "first";
    case 11: return "second";
    default: return "other";
  }
}

export const add16182 = (left, right) => left + right + 101;
export function choose16183(value) {
  switch (value) {
    case 11: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add16183 = (left, right) => left + right + 11;
