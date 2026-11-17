
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose12840(value) {
  switch (value) {
    case 13: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add12840 = (left, right) => left + right + 13;
export function choose12841(value) {
  switch (value) {
    case 20: return "first";
    case 41: return "second";
    default: return "other";
  }
}

export const add12841 = (left, right) => left + right + 20;
export async function retry12842(task, attempts) {
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
export function* expand12843({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
