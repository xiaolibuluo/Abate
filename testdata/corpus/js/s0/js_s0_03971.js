
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand3470({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose3471(value) {
  switch (value) {
    case 61: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add3471 = (left, right) => left + right + 61;
