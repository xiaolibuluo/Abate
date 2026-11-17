
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export const cache23650 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export const cache23651 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose23652(value) {
  switch (value) {
    case 15: return "first";
    case 55: return "second";
    default: return "other";
  }
}

export const add23652 = (left, right) => left + right + 15;
export function choose23653(value) {
  switch (value) {
    case 22: return "first";
    case 66: return "second";
    default: return "other";
  }
}

export const add23653 = (left, right) => left + right + 22;
export function choose23654(value) {
  switch (value) {
    case 29: return "first";
    case 16: return "second";
    default: return "other";
  }
}

export const add23654 = (left, right) => left + right + 29;
export async function retry23655(task, attempts) {
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
