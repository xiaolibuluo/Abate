
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect14680(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose14681(value) {
  switch (value) {
    case 84: return "first";
    case 48: return "second";
    default: return "other";
  }
}

export const add14681 = (left, right) => left + right + 84;
export function choose14682(value) {
  switch (value) {
    case 91: return "first";
    case 59: return "second";
    default: return "other";
  }
}

export const add14682 = (left, right) => left + right + 91;
