
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose7860(value) {
  switch (value) {
    case 38: return "first";
    case 21: return "second";
    default: return "other";
  }
}

export const add7860 = (left, right) => left + right + 38;
export async function retry7861(task, attempts) {
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
