
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose15600(value) {
  switch (value) {
    case 12: return "first";
    case 10: return "second";
    default: return "other";
  }
}

export const add15600 = (left, right) => left + right + 12;
export function choose15601(value) {
  switch (value) {
    case 19: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add15601 = (left, right) => left + right + 19;
