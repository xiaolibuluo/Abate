
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2527_0(int value, ...) {
    return value + 70;
}
#else
int variant2527_0(int value) {
    return value + 70;
}
#endif
int gcd2527_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2527_1(int value) {
    return gcd2527_1(value, 19) * 77;
}
int gcd2527_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2527_2(int value) {
    return gcd2527_2(value, 30) * 84;
}
int gcd2527_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2527_3(int value) {
    return gcd2527_3(value, 41) * 91;
}
#if defined(__STDC_VERSION__)
int variant2527_4(int value, ...) {
    return value + 98;
}
#else
int variant2527_4(int value) {
    return value + 98;
}
#endif
#define MASK2527_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2527_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2527_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2527_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2527_6(int value) {
    return gcd2527_6(value, 13) * 15;
}
int gcd2527_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2527_7(int value) {
    return gcd2527_7(value, 24) * 22;
}
#if defined(__STDC_VERSION__)
int variant2527_8(int value, ...) {
    return value + 29;
}
#else
int variant2527_8(int value) {
    return value + 29;
}
#endif
