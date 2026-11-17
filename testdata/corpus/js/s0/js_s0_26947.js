
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose23840(value) {
  switch (value) {
    case 54: return "first";
    case 44: return "second";
    default: return "other";
  }
}

export const add23840 = (left, right) => left + right + 54;
