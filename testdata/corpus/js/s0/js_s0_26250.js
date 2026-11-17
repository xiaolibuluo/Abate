
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect23210(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export const cache23211 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export async function collect23212(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose23213(value) {
  switch (value) {
    case 32: return "first";
    case 63: return "second";
    default: return "other";
  }
}

export const add23213 = (left, right) => left + right + 32;
export async function retry23214(task, attempts) {
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
export function choose23215(value) {
  switch (value) {
    case 46: return "first";
    case 24: return "second";
    default: return "other";
  }
}

export const add23215 = (left, right) => left + right + 46;
export function choose23216(value) {
  switch (value) {
    case 53: return "first";
    case 35: return "second";
    default: return "other";
  }
}

export const add23216 = (left, right) => left + right + 53;
