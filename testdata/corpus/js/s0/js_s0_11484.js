
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose10360(value) {
  switch (value) {
    case 87: return "first";
    case 63: return "second";
    default: return "other";
  }
}

export const add10360 = (left, right) => left + right + 87;
export async function retry10361(task, attempts) {
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
export function choose10362(value) {
  switch (value) {
    case 101: return "first";
    case 24: return "second";
    default: return "other";
  }
}

export const add10362 = (left, right) => left + right + 101;
