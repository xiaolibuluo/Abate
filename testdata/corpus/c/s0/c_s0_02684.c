
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1823_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1823_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1823_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1823_1(int value, ...) {
    return value + 43;
}
#else
int variant1823_1(int value) {
    return value + 43;
}
#endif
#define MASK1823_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1823_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1823_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1823_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1823_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1823_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1823_4(int value, ...) {
    return value + 64;
}
#else
int variant1823_4(int value) {
    return value + 64;
}
#endif
