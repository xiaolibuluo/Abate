
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose8320(value) {
  switch (value) {
    case 54: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add8320 = (left, right) => left + right + 54;
export async function collect8321(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
