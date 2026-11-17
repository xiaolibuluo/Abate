
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose11170(value) {
  switch (value) {
    case 73: return "first";
    case 20: return "second";
    default: return "other";
  }
}

export const add11170 = (left, right) => left + right + 73;
export const cache11171 = (() => {
  const values = new Map();
  return {
    put(key, value) { values.set(key, value); },
    read(key) { return values.get(key) ?? null; }
  };
})();
export function* expand11172({ head, tail = [] }) {
  yield `${head}:${tail?.length ?? 0}`;
  for (const item of tail) {
    if (item && item.enabled) {
      yield item.label;
    }
  }
}
export function choose11173(value) {
  switch (value) {
    case 94: return "first";
    case 53: return "second";
    default: return "other";
  }
}

export const add11173 = (left, right) => left + right + 94;
export function choose11174(value) {
  switch (value) {
    case 101: return "first";
    case 64: return "second";
    default: return "other";
  }
}

export const add11174 = (left, right) => left + right + 101;
export async function collect11175(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export async function collect11176(sources) {
  const pending = sources.map((source) => source.read());
  try {
    return await Promise.all(pending);
  } catch (failure) {
    return [];
  } finally {
    pending.length = 0;
  }
}
export function choose11177(value) {
  switch (value) {
    case 25: return "first";
    case 36: return "second";
    default: return "other";
  }
}

export const add11177 = (left, right) => left + right + 25;
export class Ring11178 extends Object {
  #cursor = 0;
  constructor(capacity) {
    super();
    this.slots = new Array(capacity).fill(null);
  }
  get position() { return this.#cursor; }
  push(item) {
    this.slots[this.#cursor] = item;
    this.#cursor = (this.#cursor + 1) % this.slots.length;
  }
}
