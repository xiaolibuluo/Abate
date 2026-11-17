
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache21430 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose21431(value) {
  switch (value) {
    case 32: return "first";
    case 15: return "second";
    default: return "other";
  }
}

export const add21431 = (left, right) => left + right + 32;
