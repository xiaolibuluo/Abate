
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1723_0(int value, ...) {
    return value + 94;
}
#else
int variant1723_0(int value) {
    return value + 94;
}
#endif
#define MASK1723_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1723_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1723_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1723_2(int value, ...) {
    return value + 11;
}
#else
int variant1723_2(int value) {
    return value + 11;
}
#endif
#if defined(__STDC_VERSION__)
int variant1723_3(int value, ...) {
    return value + 18;
}
#else
int variant1723_3(int value) {
    return value + 18;
}
#endif
#define MASK1723_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1723_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1723_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
