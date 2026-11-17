
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose9250(value) {
  switch (value) {
    case 99: return "first";
    case 18: return "second";
    default: return "other";
  }
}

export const add9250 = (left, right) => left + right + 99;
