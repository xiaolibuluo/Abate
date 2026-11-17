
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache23410 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose23411(value) {
  switch (value) {
    case 84: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add23411 = (left, right) => left + right + 84;
