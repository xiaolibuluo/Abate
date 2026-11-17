
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose4900(value) {
  switch (value) {
    case 70: return "first";
    case 23: return "second";
    default: return "other";
  }
}

export const add4900 = (left, right) => left + right + 70;
export async function collect4901(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
