
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose9430(value) {
  switch (value) {
    case 42: return "first";
    case 22: return "second";
    default: return "other";
  }
}

export const add9430 = (left, right) => left + right + 42;
export function choose9431(value) {
  switch (value) {
    case 49: return "first";
    case 33: return "second";
    default: return "other";
  }
}

export const add9431 = (left, right) => left + right + 49;
export function* expand9432({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export async function retry9433(task, attempts) {
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
