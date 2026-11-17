
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache6850 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose6851(value) {
  switch (value) {
    case 90: return "first";
    case 57: return "second";
    default: return "other";
  }
}

export const add6851 = (left, right) => left + right + 90;
