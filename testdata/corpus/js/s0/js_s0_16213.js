
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose14370(value) {
  switch (value) {
    case 62: return "first";
    case 64: return "second";
    default: return "other";
  }
}

export const add14370 = (left, right) => left + right + 62;
export async function retry14371(task, attempts) {
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
export async function retry14372(task, attempts) {
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
export function* expand14373({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose14374(value) {
  switch (value) {
    case 90: return "first";
    case 47: return "second";
    default: return "other";
  }
}

export const add14374 = (left, right) => left + right + 90;
