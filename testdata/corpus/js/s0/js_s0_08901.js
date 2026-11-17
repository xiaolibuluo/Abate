
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry7650(task, attempts) {
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
export async function retry7651(task, attempts) {
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
export async function collect7652(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose7653(value) {
  switch (value) {
    case 77: return "first";
    case 29: return "second";
    default: return "other";
  }
}

export const add7653 = (left, right) => left + right + 77;
export const cache7654 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose7655(value) {
  switch (value) {
    case 91: return "first";
    case 51: return "second";
    default: return "other";
  }
}

export const add7655 = (left, right) => left + right + 91;
export function choose7656(value) {
  switch (value) {
    case 98: return "first";
    case 62: return "second";
    default: return "other";
  }
}

export const add7656 = (left, right) => left + right + 98;
