
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache13150 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose13151(value) {
  switch (value) {
    case 35: return "first";
    case 14: return "second";
    default: return "other";
  }
}

export const add13151 = (left, right) => left + right + 35;
