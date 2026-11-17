
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose9760(value) {
  switch (value) {
    case 83: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add9760 = (left, right) => left + right + 83;
export function* expand9761({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
