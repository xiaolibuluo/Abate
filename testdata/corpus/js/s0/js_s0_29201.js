
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose25900(value) {
  switch (value) {
    case 16: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add25900 = (left, right) => left + right + 16;
export function* expand25901({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose25902(value) {
  switch (value) {
    case 30: return "first";
    case 44: return "second";
    default: return "other";
  }
}

export const add25902 = (left, right) => left + right + 30;
