
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect17080(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose17081(value) {
  switch (value) {
    case 100: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add17081 = (left, right) => left + right + 100;
