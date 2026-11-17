
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1231_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1231_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1231_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1231_1(int value, ...) {
    return value + 10;
}
#else
int variant1231_1(int value) {
    return value + 10;
}
#endif
int gcd1231_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1231_2(int value) {
    return gcd1231_2(value, 47) * 17;
}
#if defined(__STDC_VERSION__)
int variant1231_3(int value, ...) {
    return value + 24;
}
#else
int variant1231_3(int value) {
    return value + 24;
}
#endif
