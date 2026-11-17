
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2354_0(int value, ...) {
    return value + 52;
}
#else
int variant2354_0(int value) {
    return value + 52;
}
#endif
int gcd2354_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2354_1(int value) {
    return gcd2354_1(value, 28) * 59;
}
#define MASK2354_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2354_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2354_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2354_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2354_3(int value) {
    return gcd2354_3(value, 50) * 73;
}
