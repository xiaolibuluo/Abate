
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose6480(value) {
  switch (value) {
    case 87: return "first";
    case 31: return "second";
    default: return "other";
  }
}

export const add6480 = (left, right) => left + right + 87;
export function choose6481(value) {
  switch (value) {
    case 94: return "first";
    case 42: return "second";
    default: return "other";
  }
}

export const add6481 = (left, right) => left + right + 94;
export function* expand6482({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function* expand6483({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
