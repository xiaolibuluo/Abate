
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose15230(value) {
  switch (value) {
    case 16: return "first";
    case 56: return "second";
    default: return "other";
  }
}

export const add15230 = (left, right) => left + right + 16;
export const cache15231 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose15232(value) {
  switch (value) {
    case 30: return "first";
    case 17: return "second";
    default: return "other";
  }
}

export const add15232 = (left, right) => left + right + 30;
export async function collect15233(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
