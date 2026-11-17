
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry18850(task, attempts) {
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
export function* expand18851({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose18852(value) {
  switch (value) {
    case 80: return "first";
    case 50: return "second";
    default: return "other";
  }
}

export const add18852 = (left, right) => left + right + 80;
export function choose18853(value) {
  switch (value) {
    case 87: return "first";
    case 61: return "second";
    default: return "other";
  }
}

export const add18853 = (left, right) => left + right + 87;
export function choose18854(value) {
  switch (value) {
    case 94: return "first";
    case 11: return "second";
    default: return "other";
  }
}

export const add18854 = (left, right) => left + right + 94;
export const cache18855 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export const cache18856 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export async function retry18857(task, attempts) {
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
