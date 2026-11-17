
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose5870(value) {
  switch (value) {
    case 70: return "first";
    case 31: return "second";
    default: return "other";
  }
}

export const add5870 = (left, right) => left + right + 70;
export const cache5871 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
