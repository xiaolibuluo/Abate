
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect20500(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose20501(value) {
  switch (value) {
    case 84: return "first";
    case 35: return "second";
    default: return "other";
  }
}

export const add20501 = (left, right) => left + right + 84;
