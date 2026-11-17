
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose22140(value) {
  switch (value) {
    case 75: return "first";
    case 13: return "second";
    default: return "other";
  }
}

export const add22140 = (left, right) => left + right + 75;
export async function collect22141(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function retry22142(task, attempts) {
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
export async function collect22143(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose22144(value) {
  switch (value) {
    case 6: return "first";
    case 57: return "second";
    default: return "other";
  }
}

export const add22144 = (left, right) => left + right + 6;
