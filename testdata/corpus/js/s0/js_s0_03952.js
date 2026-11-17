
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose3420(value) {
  switch (value) {
    case 86: return "first";
    case 24: return "second";
    default: return "other";
  }
}

export const add3420 = (left, right) => left + right + 86;
