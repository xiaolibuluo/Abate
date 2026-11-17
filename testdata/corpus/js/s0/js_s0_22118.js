
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect19370(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function retry19371(task, attempts) {
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
export function choose19372(value) {
  switch (value) {
    case 77: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add19372 = (left, right) => left + right + 77;
export const cache19373 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose19374(value) {
  switch (value) {
    case 91: return "first";
    case 9: return "second";
    default: return "other";
  }
}

export const add19374 = (left, right) => left + right + 91;
export function choose19375(value) {
  switch (value) {
    case 98: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add19375 = (left, right) => left + right + 98;
export async function collect19376(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose19377(value) {
  switch (value) {
    case 15: return "first";
    case 42: return "second";
    default: return "other";
  }
}

export const add19377 = (left, right) => left + right + 15;
