
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose14610(value) {
  switch (value) {
    case 83: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add14610 = (left, right) => left + right + 83;
