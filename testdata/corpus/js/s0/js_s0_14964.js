
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache13200 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose13201(value) {
  switch (value) {
    case 100: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add13201 = (left, right) => left + right + 100;
