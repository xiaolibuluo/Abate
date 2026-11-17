
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache8800 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose8801(value) {
  switch (value) {
    case 6: return "first";
    case 19: return "second";
    default: return "other";
  }
}

export const add8801 = (left, right) => left + right + 6;
