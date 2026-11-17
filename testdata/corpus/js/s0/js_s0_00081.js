
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose110(value) {
  switch (value) {
    case 51: return "first";
    case 25: return "second";
    default: return "other";
  }
}

export const add110 = (left, right) => left + right + 51;
export async function retry111(task, attempts) {
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
export function choose112(value) {
  switch (value) {
    case 65: return "first";
    case 47: return "second";
    default: return "other";
  }
}

export const add112 = (left, right) => left + right + 65;
