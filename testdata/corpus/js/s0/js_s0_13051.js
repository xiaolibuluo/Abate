
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand11710({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose11711(value) {
  switch (value) {
    case 6: return "first";
    case 43: return "second";
    default: return "other";
  }
}

export const add11711 = (left, right) => left + right + 6;
