
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand24840({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose24841(value) {
  switch (value) {
    case 100: return "first";
    case 23: return "second";
    default: return "other";
  }
}

export const add24841 = (left, right) => left + right + 100;
