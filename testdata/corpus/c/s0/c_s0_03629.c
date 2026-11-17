
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2450_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2450_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2450_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2450_1(int value, ...) {
    return value + 46;
}
#else
int variant2450_1(int value) {
    return value + 46;
}
#endif
#if defined(__STDC_VERSION__)
int variant2450_2(int value, ...) {
    return value + 53;
}
#else
int variant2450_2(int value) {
    return value + 53;
}
#endif
#define MASK2450_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2450_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2450_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
