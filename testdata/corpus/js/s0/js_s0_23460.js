
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand20630({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export const cache20631 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose20632(value) {
  switch (value) {
    case 66: return "first";
    case 15: return "second";
    default: return "other";
  }
}

export const add20632 = (left, right) => left + right + 66;
export function choose20633(value) {
  switch (value) {
    case 73: return "first";
    case 26: return "second";
    default: return "other";
  }
}

export const add20633 = (left, right) => left + right + 73;
export function* expand20634({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
