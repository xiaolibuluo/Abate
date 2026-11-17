
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2252_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2252_0(int value) {
    return gcd2252_0(value, 35) * 84;
}
#if defined(__STDC_VERSION__)
int variant2252_1(int value, ...) {
    return value + 91;
}
#else
int variant2252_1(int value) {
    return value + 91;
}
#endif
#if defined(__STDC_VERSION__)
int variant2252_2(int value, ...) {
    return value + 98;
}
#else
int variant2252_2(int value) {
    return value + 98;
}
#endif
#define MASK2252_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2252_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2252_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
