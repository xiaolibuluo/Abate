
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1707_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1707_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1707_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1707_1(int value, ...) {
    return value + 87;
}
#else
int variant1707_1(int value) {
    return value + 87;
}
#endif
