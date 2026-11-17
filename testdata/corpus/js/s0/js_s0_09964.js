
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose8820(value) {
  switch (value) {
    case 25: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add8820 = (left, right) => left + right + 25;
export const cache8821 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
