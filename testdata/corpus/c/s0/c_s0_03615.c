
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2443_0(int value, ...) {
    return value + 45;
}
#else
int variant2443_0(int value) {
    return value + 45;
}
#endif
#if defined(__STDC_VERSION__)
int variant2443_1(int value, ...) {
    return value + 52;
}
#else
int variant2443_1(int value) {
    return value + 52;
}
#endif
#define MASK2443_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2443_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2443_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
