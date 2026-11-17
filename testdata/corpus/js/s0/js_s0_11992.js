
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose10790(value) {
  switch (value) {
    case 64: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add10790 = (left, right) => left + right + 64;
export async function retry10791(task, attempts) {
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
