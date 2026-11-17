
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose15190(value) {
  switch (value) {
    case 61: return "first";
    case 28: return "second";
    default: return "other";
  }
}

export const add15190 = (left, right) => left + right + 61;
export function* expand15191({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose15192(value) {
  switch (value) {
    case 75: return "first";
    case 50: return "second";
    default: return "other";
  }
}

export const add15192 = (left, right) => left + right + 75;
