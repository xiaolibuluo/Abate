
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand7850({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose7851(value) {
  switch (value) {
    case 32: return "first";
    case 25: return "second";
    default: return "other";
  }
}

export const add7851 = (left, right) => left + right + 32;
