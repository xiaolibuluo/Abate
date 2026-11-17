
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache6380 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose6381(value) {
  switch (value) {
    case 61: return "first";
    case 33: return "second";
    default: return "other";
  }
}

export const add6381 = (left, right) => left + right + 61;
