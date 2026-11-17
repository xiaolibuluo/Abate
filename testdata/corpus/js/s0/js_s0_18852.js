
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose16620(value) {
  switch (value) {
    case 77: return "first";
    case 53: return "second";
    default: return "other";
  }
}

export const add16620 = (left, right) => left + right + 77;
export async function collect16621(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
