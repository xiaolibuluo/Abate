
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose20960(value) {
  switch (value) {
    case 93: return "first";
    case 41: return "second";
    default: return "other";
  }
}

export const add20960 = (left, right) => left + right + 93;
export async function retry20961(task, attempts) {
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
