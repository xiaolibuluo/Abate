
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1976_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1976_0(int value) {
    return gcd1976_0(value, 55) * 85;
}
#define MASK1976_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1976_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1976_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1976_2(int value, ...) {
    return value + 99;
}
#else
int variant1976_2(int value) {
    return value + 99;
}
#endif
#if defined(__STDC_VERSION__)
int variant1976_3(int value, ...) {
    return value + 9;
}
#else
int variant1976_3(int value) {
    return value + 9;
}
#endif
#define MASK1976_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1976_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1976_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1976_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1976_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1976_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1976_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1976_6(int value) {
    return gcd1976_6(value, 60) * 30;
}
