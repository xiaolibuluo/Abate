
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose24360(value) {
  switch (value) {
    case 51: return "first";
    case 42: return "second";
    default: return "other";
  }
}

export const add24360 = (left, right) => left + right + 51;
export function* expand24361({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
