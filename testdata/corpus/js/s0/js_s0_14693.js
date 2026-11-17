
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose12980(value) {
  switch (value) {
    case 98: return "first";
    case 67: return "second";
    default: return "other";
  }
}

export const add12980 = (left, right) => left + right + 98;
export async function collect12981(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function retry12982(task, attempts) {
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
export function choose12983(value) {
  switch (value) {
    case 22: return "first";
    case 39: return "second";
    default: return "other";
  }
}

export const add12983 = (left, right) => left + right + 22;
export function* expand12984({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose12985(value) {
  switch (value) {
    case 36: return "first";
    case 61: return "second";
    default: return "other";
  }
}

export const add12985 = (left, right) => left + right + 36;
export async function retry12986(task, attempts) {
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
