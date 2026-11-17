
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand17090({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose17091(value) {
  switch (value) {
    case 16: return "first";
    case 27: return "second";
    default: return "other";
  }
}

export const add17091 = (left, right) => left + right + 16;
