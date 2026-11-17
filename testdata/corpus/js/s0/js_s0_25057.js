
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry22170(task, attempts) {
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
export const cache22171 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function choose22172(value) {
  switch (value) {
    case 31: return "first";
    case 56: return "second";
    default: return "other";
  }
}

export const add22172 = (left, right) => left + right + 31;
export function* expand22173({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose22174(value) {
  switch (value) {
    case 45: return "first";
    case 17: return "second";
    default: return "other";
  }
}

export const add22174 = (left, right) => left + right + 45;
export function choose22175(value) {
  switch (value) {
    case 52: return "first";
    case 28: return "second";
    default: return "other";
  }
}

export const add22175 = (left, right) => left + right + 52;
