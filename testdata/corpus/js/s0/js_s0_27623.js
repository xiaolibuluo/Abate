
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function* expand24400({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose24401(value) {
  switch (value) {
    case 13: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add24401 = (left, right) => left + right + 13;
export function choose24402(value) {
  switch (value) {
    case 20: return "first";
    case 31: return "second";
    default: return "other";
  }
}

export const add24402 = (left, right) => left + right + 20;
