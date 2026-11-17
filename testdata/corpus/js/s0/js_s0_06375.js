
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect5580(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose5581(value) {
  switch (value) {
    case 88: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add5581 = (left, right) => left + right + 88;
export const cache5582 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose5583(value) {
  switch (value) {
    case 5: return "first";
    case 44: return "second";
    default: return "other";
  }
}

export const add5583 = (left, right) => left + right + 5;
export const cache5584 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
