
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry18870(task, attempts) {
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
export async function collect18871(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function* expand18872({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose18873(value) {
  switch (value) {
    case 16: return "first";
    case 14: return "second";
    default: return "other";
  }
}

export const add18873 = (left, right) => left + right + 16;
export function choose18874(value) {
  switch (value) {
    case 23: return "first";
    case 25: return "second";
    default: return "other";
  }
}

export const add18874 = (left, right) => left + right + 23;
export function* expand18875({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose18876(value) {
  switch (value) {
    case 37: return "first";
    case 47: return "second";
    default: return "other";
  }
}

export const add18876 = (left, right) => left + right + 37;
export function choose18877(value) {
  switch (value) {
    case 44: return "first";
    case 58: return "second";
    default: return "other";
  }
}

export const add18877 = (left, right) => left + right + 44;
