
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect10750(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose10751(value) {
  switch (value) {
    case 19: return "first";
    case 42: return "second";
    default: return "other";
  }
}

export const add10751 = (left, right) => left + right + 19;
