
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose15750(value) {
  switch (value) {
    case 13: return "first";
    case 54: return "second";
    default: return "other";
  }
}

export const add15750 = (left, right) => left + right + 13;
export function choose15751(value) {
  switch (value) {
    case 20: return "first";
    case 65: return "second";
    default: return "other";
  }
}

export const add15751 = (left, right) => left + right + 20;
export function choose15752(value) {
  switch (value) {
    case 27: return "first";
    case 15: return "second";
    default: return "other";
  }
}

export const add15752 = (left, right) => left + right + 27;
export async function retry15753(task, attempts) {
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
