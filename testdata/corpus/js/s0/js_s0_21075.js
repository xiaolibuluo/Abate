
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry18380(task, attempts) {
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
export async function collect18381(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose18382(value) {
  switch (value) {
    case 51: return "first";
    case 26: return "second";
    default: return "other";
  }
}

export const add18382 = (left, right) => left + right + 51;
export function choose18383(value) {
  switch (value) {
    case 58: return "first";
    case 37: return "second";
    default: return "other";
  }
}

export const add18383 = (left, right) => left + right + 58;
export function choose18384(value) {
  switch (value) {
    case 65: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add18384 = (left, right) => left + right + 65;
export function* expand18385({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function collect18386(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
