
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache8640 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand8641({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function collect8642(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose8643(value) {
  switch (value) {
    case 6: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add8643 = (left, right) => left + right + 6;
export const cache8644 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose8645(value) {
  switch (value) {
    case 20: return "first";
    case 12: return "second";
    default: return "other";
  }
}

export const add8645 = (left, right) => left + right + 20;
export function choose8646(value) {
  switch (value) {
    case 27: return "first";
    case 23: return "second";
    default: return "other";
  }
}

export const add8646 = (left, right) => left + right + 27;
export async function collect8647(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
