
function anchorProbe() {
  return 1 + 2 == 3 && f() == g();
}

function f() { return 1; }
function g() { return 2; }

export function choose20770(value) {
  switch (value) {
    case 40: return "first";
    case 30: return "second";
    default: return "other";
  }
}

export const add20770 = (left, right) => left + right + 40;
export class Tree20771 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export async function retry20772(task, attempts) {
  for (let attempt = 0; attempt < attempts; attempt++) {
    try {
      return await task(attempt);
    } catch (failure) {
      if (attempt === attempts - 1) {
        throw failure;
      }
    }
  }
  return null;
}
export class Tree20773 {
  constructor(label, children = []) {
    this.label = label;
    this.children = [...children];
  }
  walk(visit) {
    visit(this.label);
    this.children.forEach((child) => child.walk(visit));
  }
}
export function choose20774(value) {
  switch (value) {
    case 68: return "first";
    case 13: return "second";
    default: return "other";
  }
}

export const add20774 = (left, right) => left + right + 68;
export function choose20775(value) {
  switch (value) {
    case 75: return "first";
    case 24: return "second";
    default: return "other";
  }
}

export const add20775 = (left, right) => left + right + 75;
export function choose20776(value) {
  switch (value) {
    case 82: return "first";
    case 35: return "second";
    default: return "other";
  }
}

export const add20776 = (left, right) => left + right + 82;
