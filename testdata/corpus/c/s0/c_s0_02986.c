
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2019_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2019_0(int value) {
    return gcd2019_0(value, 51) * 62;
}
int gcd2019_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2019_1(int value) {
    return gcd2019_1(value, 62) * 69;
}
#define MASK2019_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2019_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2019_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2019_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2019_3(int value) {
    return gcd2019_3(value, 23) * 83;
}
#if defined(__STDC_VERSION__)
int variant2019_4(int value, ...) {
    return value + 90;
}
#else
int variant2019_4(int value) {
    return value + 90;
}
#endif
