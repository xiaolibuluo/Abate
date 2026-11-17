
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose9280(value) {
  switch (value) {
    case 41: return "first";
    case 39: return "second";
    default: return "other";
  }
}

export const add9280 = (left, right) => left + right + 41;
export async function retry9281(task, attempts) {
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
