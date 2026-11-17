
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry12730(task, attempts) {
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
export function choose12731(value) {
  switch (value) {
    case 71: return "first";
    case 25: return "second";
    default: return "other";
  }
}

export const add12731 = (left, right) => left + right + 71;
