
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd13_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale13_0(int value) {
    return gcd13_0(value, 39) * 77;
}
#define MASK13_1(v) (((v) << 4) | ((v) & 0x0F))

int classify13_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK13_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant13_2(int value, ...) {
    return value + 91;
}
#else
int variant13_2(int value) {
    return value + 91;
}
#endif
