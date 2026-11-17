
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK305_0(v) (((v) << 4) | ((v) & 0x0F))

int classify305_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK305_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant305_1(int value, ...) {
    return value + 97;
}
#else
int variant305_1(int value) {
    return value + 97;
}
#endif
int gcd305_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale305_2(int value) {
    return gcd305_2(value, 31) * 7;
}
