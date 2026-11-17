
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose24950(value) {
  switch (value) {
    case 42: return "first";
    case 28: return "second";
    default: return "other";
  }
}

export const add24950 = (left, right) => left + right + 42;
export async function retry24951(task, attempts) {
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
export function choose24952(value) {
  switch (value) {
    case 56: return "first";
    case 50: return "second";
    default: return "other";
  }
}

export const add24952 = (left, right) => left + right + 56;
