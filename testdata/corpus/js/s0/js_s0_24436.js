
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry21520(task, attempts) {
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
export function choose21521(value) {
  switch (value) {
    case 52: return "first";
    case 17: return "second";
    default: return "other";
  }
}

export const add21521 = (left, right) => left + right + 52;
export function choose21522(value) {
  switch (value) {
    case 59: return "first";
    case 28: return "second";
    default: return "other";
  }
}

export const add21522 = (left, right) => left + right + 59;
