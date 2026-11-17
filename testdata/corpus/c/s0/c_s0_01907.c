
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1284_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1284_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1284_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1284_1(int value, ...) {
    return value + 20;
}
#else
int variant1284_1(int value) {
    return value + 20;
}
#endif
int gcd1284_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1284_2(int value) {
    return gcd1284_2(value, 52) * 27;
}
#if defined(__STDC_VERSION__)
int variant1284_3(int value, ...) {
    return value + 34;
}
#else
int variant1284_3(int value) {
    return value + 34;
}
#endif
#define MASK1284_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1284_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1284_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
