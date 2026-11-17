
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2014_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2014_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2014_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2014_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2014_1(int value) {
    return gcd2014_1(value, 27) * 101;
}
#if defined(__STDC_VERSION__)
int variant2014_2(int value, ...) {
    return value + 11;
}
#else
int variant2014_2(int value) {
    return value + 11;
}
#endif
int gcd2014_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2014_3(int value) {
    return gcd2014_3(value, 49) * 18;
}
