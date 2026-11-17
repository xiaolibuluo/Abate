
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache16770 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export async function collect16771(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose16772(value) {
  switch (value) {
    case 92: return "first";
    case 58: return "second";
    default: return "other";
  }
}

export const add16772 = (left, right) => left + right + 92;
export const cache16773 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose16774(value) {
  switch (value) {
    case 9: return "first";
    case 19: return "second";
    default: return "other";
  }
}

export const add16774 = (left, right) => left + right + 9;
