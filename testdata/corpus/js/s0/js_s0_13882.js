
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose12300(value) {
  switch (value) {
    case 87: return "first";
    case 18: return "second";
    default: return "other";
  }
}

export const add12300 = (left, right) => left + right + 87;
export function choose12301(value) {
  switch (value) {
    case 94: return "first";
    case 29: return "second";
    default: return "other";
  }
}

export const add12301 = (left, right) => left + right + 94;
export async function retry12302(task, attempts) {
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
