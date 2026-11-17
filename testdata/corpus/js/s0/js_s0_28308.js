
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache25080 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose25081(value) {
  switch (value) {
    case 24: return "first";
    case 69: return "second";
    default: return "other";
  }
}

export const add25081 = (left, right) => left + right + 24;
export function choose25082(value) {
  switch (value) {
    case 31: return "first";
    case 19: return "second";
    default: return "other";
  }
}

export const add25082 = (left, right) => left + right + 31;
export function choose25083(value) {
  switch (value) {
    case 38: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add25083 = (left, right) => left + right + 38;
export async function retry25084(task, attempts) {
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
export async function collect25085(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
