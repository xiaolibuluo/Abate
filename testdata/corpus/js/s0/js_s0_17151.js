
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose15140(value) {
  switch (value) {
    case 93: return "first";
    case 54: return "second";
    default: return "other";
  }
}

export const add15140 = (left, right) => left + right + 93;
export async function retry15141(task, attempts) {
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
