
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2010_0(int value, ...) {
    return value + 42;
}
#else
int variant2010_0(int value) {
    return value + 42;
}
#endif
#define MASK2010_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2010_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2010_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2010_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2010_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2010_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2010_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2010_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2010_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
