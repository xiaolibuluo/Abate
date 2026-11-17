
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2268_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2268_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2268_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2268_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2268_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2268_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2268_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2268_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2268_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2268_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2268_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2268_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2268_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2268_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2268_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2268_5(int value, ...) {
    return value + 36;
}
#else
int variant2268_5(int value) {
    return value + 36;
}
#endif
