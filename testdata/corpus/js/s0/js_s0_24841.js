
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose21910(value) {
  switch (value) {
    case 67: return "first";
    case 35: return "second";
    default: return "other";
  }
}

export const add21910 = (left, right) => left + right + 67;
