
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose21440(value) {
  switch (value) {
    case 38: return "first";
    case 11: return "second";
    default: return "other";
  }
}

export const add21440 = (left, right) => left + right + 38;
export async function retry21441(task, attempts) {
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
