
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand16100({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose16101(value) {
  switch (value) {
    case 87: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add16101 = (left, right) => left + right + 87;
