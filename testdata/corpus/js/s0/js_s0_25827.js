
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand22880({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose22881(value) {
  switch (value) {
    case 74: return "first";
    case 54: return "second";
    default: return "other";
  }
}

export const add22881 = (left, right) => left + right + 74;
