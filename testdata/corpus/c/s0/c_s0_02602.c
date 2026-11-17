
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1765_0(int value, ...) {
    return value + 58;
}
#else
int variant1765_0(int value) {
    return value + 58;
}
#endif
int gcd1765_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1765_1(int value) {
    return gcd1765_1(value, 53) * 65;
}
#if defined(__STDC_VERSION__)
int variant1765_2(int value, ...) {
    return value + 72;
}
#else
int variant1765_2(int value) {
    return value + 72;
}
#endif
#define MASK1765_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1765_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1765_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
