
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function retry24480(task, attempts) {
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
export function choose24481(value) {
  switch (value) {
    case 20: return "first";
    case 15: return "second";
    default: return "other";
  }
}

export const add24481 = (left, right) => left + right + 20;
export function choose24482(value) {
  switch (value) {
    case 27: return "first";
    case 26: return "second";
    default: return "other";
  }
}

export const add24482 = (left, right) => left + right + 27;
export function choose24483(value) {
  switch (value) {
    case 34: return "first";
    case 37: return "second";
    default: return "other";
  }
}

export const add24483 = (left, right) => left + right + 34;
