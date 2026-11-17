
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose3920(value) {
  switch (value) {
    case 57: return "first";
    case 69: return "second";
    default: return "other";
  }
}

export const add3920 = (left, right) => left + right + 57;
export async function collect3921(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
