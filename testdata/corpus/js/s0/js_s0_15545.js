
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect13790(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose13791(value) {
  switch (value) {
    case 91: return "first";
    case 35: return "second";
    default: return "other";
  }
}

export const add13791 = (left, right) => left + right + 91;
export async function retry13792(task, attempts) {
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
export function choose13793(value) {
  switch (value) {
    case 8: return "first";
    case 57: return "second";
    default: return "other";
  }
}

export const add13793 = (left, right) => left + right + 8;
