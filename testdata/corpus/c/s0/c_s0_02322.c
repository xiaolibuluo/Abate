
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1566_0(int value, ...) {
    return value + 90;
}
#else
int variant1566_0(int value) {
    return value + 90;
}
#endif
int gcd1566_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1566_1(int value) {
    return gcd1566_1(value, 63) * 97;
}
#define MASK1566_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1566_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1566_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
