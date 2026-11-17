
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1229_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1229_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1229_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1229_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1229_1(int value) {
    return gcd1229_1(value, 22) * 81;
}
#if defined(__STDC_VERSION__)
int variant1229_2(int value, ...) {
    return value + 88;
}
#else
int variant1229_2(int value) {
    return value + 88;
}
#endif
