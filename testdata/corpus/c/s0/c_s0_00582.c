
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant398_0(int value, ...) {
    return value + 38;
}
#else
int variant398_0(int value) {
    return value + 38;
}
#endif
int gcd398_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale398_1(int value) {
    return gcd398_1(value, 61) * 45;
}
#define MASK398_2(v) (((v) << 4) | ((v) & 0x0F))

int classify398_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK398_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
