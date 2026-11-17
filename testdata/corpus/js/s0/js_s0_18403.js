
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect16210(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export const cache16211 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose16212(value) {
  switch (value) {
    case 43: return "first";
    case 32: return "second";
    default: return "other";
  }
}

export const add16212 = (left, right) => left + right + 43;
export function choose16213(value) {
  switch (value) {
    case 50: return "first";
    case 43: return "second";
    default: return "other";
  }
}

export const add16213 = (left, right) => left + right + 50;
