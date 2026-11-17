
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose8840(value) {
  switch (value) {
    case 51: return "first";
    case 36: return "second";
    default: return "other";
  }
}

export const add8840 = (left, right) => left + right + 51;
export function choose8841(value) {
  switch (value) {
    case 58: return "first";
    case 47: return "second";
    default: return "other";
  }
}

export const add8841 = (left, right) => left + right + 58;
export const cache8842 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
