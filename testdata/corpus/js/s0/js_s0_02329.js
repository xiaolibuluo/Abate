
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect2080(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose2081(value) {
  switch (value) {
    case 97: return "first";
    case 12: return "second";
    default: return "other";
  }
}

export const add2081 = (left, right) => left + right + 97;
export function choose2082(value) {
  switch (value) {
    case 7: return "first";
    case 23: return "second";
    default: return "other";
  }
}

export const add2082 = (left, right) => left + right + 7;
