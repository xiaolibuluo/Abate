
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache8290 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose8291(value) {
  switch (value) {
    case 22: return "first";
    case 28: return "second";
    default: return "other";
  }
}

export const add8291 = (left, right) => left + right + 22;
