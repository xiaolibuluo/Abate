
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand12200({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose12201(value) {
  switch (value) {
    case 61: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add12201 = (left, right) => left + right + 61;
