
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose26290(value) {
  switch (value) {
    case 38: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add26290 = (left, right) => left + right + 38;
