
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose3080(value) {
  switch (value) {
    case 32: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add3080 = (left, right) => left + right + 32;
export const cache3081 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose3082(value) {
  switch (value) {
    case 46: return "first";
    case 52: return "second";
    default: return "other";
  }
}

export const add3082 = (left, right) => left + right + 46;
export async function collect3083(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
