
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1322_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1322_0(int value) {
    return gcd1322_0(value, 52) * 22;
}
#if defined(__STDC_VERSION__)
int variant1322_1(int value, ...) {
    return value + 29;
}
#else
int variant1322_1(int value) {
    return value + 29;
}
#endif
#define MASK1322_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1322_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1322_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
