
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2292_0(int value, ...) {
    return value + 22;
}
#else
int variant2292_0(int value) {
    return value + 22;
}
#endif
#define MASK2292_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2292_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2292_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
