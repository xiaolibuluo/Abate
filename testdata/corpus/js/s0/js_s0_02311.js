
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache2060 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose2061(value) {
  switch (value) {
    case 71: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add2061 = (left, right) => left + right + 71;
export function choose2062(value) {
  switch (value) {
    case 78: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add2062 = (left, right) => left + right + 78;
