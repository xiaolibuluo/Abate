
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1478_0(int value, ...) {
    return value + 13;
}
#else
int variant1478_0(int value) {
    return value + 13;
}
#endif
#define MASK1478_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1478_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1478_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1478_2(int value, ...) {
    return value + 27;
}
#else
int variant1478_2(int value) {
    return value + 27;
}
#endif
#define MASK1478_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1478_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1478_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1478_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1478_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1478_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
