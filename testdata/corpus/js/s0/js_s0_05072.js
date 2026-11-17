
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect4420(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose4421(value) {
  switch (value) {
    case 35: return "first";
    case 64: return "second";
    default: return "other";
  }
}

export const add4421 = (left, right) => left + right + 35;
