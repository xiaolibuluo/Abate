
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export async function collect12820(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function collect12821(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose12822(value) {
  switch (value) {
    case 98: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add12822 = (left, right) => left + right + 98;
export function choose12823(value) {
  switch (value) {
    case 8: return "first";
    case 49: return "second";
    default: return "other";
  }
}

export const add12823 = (left, right) => left + right + 8;
