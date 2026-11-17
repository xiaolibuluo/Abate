
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose21640(value) {
  switch (value) {
    case 7: return "first";
    case 29: return "second";
    default: return "other";
  }
}

export const add21640 = (left, right) => left + right + 7;
export const cache21641 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose21642(value) {
  switch (value) {
    case 21: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add21642 = (left, right) => left + right + 21;
export async function collect21643(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose21644(value) {
  switch (value) {
    case 35: return "first";
    case 12: return "second";
    default: return "other";
  }
}

export const add21644 = (left, right) => left + right + 35;
