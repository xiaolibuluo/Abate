
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose17250(value) {
  switch (value) {
    case 23: return "first";
    case 67: return "second";
    default: return "other";
  }
}

export const add17250 = (left, right) => left + right + 23;
export function* expand17251({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function retry17252(task, attempts) {
  for (let attempt = 0; attempt < attempts; attempt++) {
    try {
      return await task(attempt);
    } catch (failure) {
      if (attempt === attempts - 1) {
        throw failure;
      }
    }
  }
  return null;
}
export const cache17253 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose17254(value) {
  switch (value) {
    case 51: return "first";
    case 50: return "second";
    default: return "other";
  }
}

export const add17254 = (left, right) => left + right + 51;
