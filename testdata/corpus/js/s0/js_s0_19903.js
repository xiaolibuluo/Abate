
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache17510 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose17511(value) {
  switch (value) {
    case 77: return "first";
    case 16: return "second";
    default: return "other";
  }
}

export const add17511 = (left, right) => left + right + 77;
export function choose17512(value) {
  switch (value) {
    case 84: return "first";
    case 27: return "second";
    default: return "other";
  }
}

export const add17512 = (left, right) => left + right + 84;
export async function retry17513(task, attempts) {
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
export async function collect17514(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function retry17515(task, attempts) {
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
export function choose17516(value) {
  switch (value) {
    case 15: return "first";
    case 10: return "second";
    default: return "other";
  }
}

export const add17516 = (left, right) => left + right + 15;
export async function collect17517(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export const cache17518 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
