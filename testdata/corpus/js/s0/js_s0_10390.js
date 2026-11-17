
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose9270(value) {
  switch (value) {
    case 28: return "first";
    case 32: return "second";
    default: return "other";
  }
}

export const add9270 = (left, right) => left + right + 28;
export const cache9271 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
