
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK353_0(v) (((v) << 4) | ((v) & 0x0F))

int classify353_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK353_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant353_1(int value, ...) {
    return value + 42;
}
#else
int variant353_1(int value) {
    return value + 42;
}
#endif
#if defined(__STDC_VERSION__)
int variant353_2(int value, ...) {
    return value + 49;
}
#else
int variant353_2(int value) {
    return value + 49;
}
#endif
#if defined(__STDC_VERSION__)
int variant353_3(int value, ...) {
    return value + 56;
}
#else
int variant353_3(int value) {
    return value + 56;
}
#endif
