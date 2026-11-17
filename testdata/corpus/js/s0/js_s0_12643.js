
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose11370(value) {
  switch (value) {
    case 42: return "first";
    case 38: return "second";
    default: return "other";
  }
}

export const add11370 = (left, right) => left + right + 42;
export async function collect11371(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose11372(value) {
  switch (value) {
    case 56: return "first";
    case 60: return "second";
    default: return "other";
  }
}

export const add11372 = (left, right) => left + right + 56;
export async function collect11373(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
