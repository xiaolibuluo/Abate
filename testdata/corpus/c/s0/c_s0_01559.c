
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1053_0(int value, ...) {
    return value + 17;
}
#else
int variant1053_0(int value) {
    return value + 17;
}
#endif
int gcd1053_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1053_1(int value) {
    return gcd1053_1(value, 10) * 24;
}
int gcd1053_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1053_2(int value) {
    return gcd1053_2(value, 21) * 31;
}
#define MASK1053_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1053_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1053_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1053_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1053_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1053_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1053_5(int value, ...) {
    return value + 52;
}
#else
int variant1053_5(int value) {
    return value + 52;
}
#endif
#define MASK1053_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1053_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1053_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
