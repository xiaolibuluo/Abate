
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose19500(value) {
  switch (value) {
    case 38: return "first";
    case 56: return "second";
    default: return "other";
  }
}

export const add19500 = (left, right) => left + right + 38;
export function choose19501(value) {
  switch (value) {
    case 45: return "first";
    case 67: return "second";
    default: return "other";
  }
}

export const add19501 = (left, right) => left + right + 45;
