
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect10280(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose10281(value) {
  switch (value) {
    case 87: return "first";
    case 18: return "second";
    default: return "other";
  }
}

export const add10281 = (left, right) => left + right + 87;
