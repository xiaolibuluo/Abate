
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry17260(task, attempts) {
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
export async function retry17261(task, attempts) {
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
export function choose17262(value) {
  switch (value) {
    case 50: return "first";
    case 35: return "second";
    default: return "other";
  }
}

export const add17262 = (left, right) => left + right + 50;
export function* expand17263({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose17264(value) {
  switch (value) {
    case 64: return "first";
    case 57: return "second";
    default: return "other";
  }
}

export const add17264 = (left, right) => left + right + 64;
