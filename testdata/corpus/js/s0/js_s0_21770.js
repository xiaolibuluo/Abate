
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose18990(value) {
  switch (value) {
    case 54: return "first";
    case 65: return "second";
    default: return "other";
  }
}

export const add18990 = (left, right) => left + right + 54;
