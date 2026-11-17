
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2122_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2122_0(int value) {
    return gcd2122_0(value, 40) * 43;
}
#if defined(__STDC_VERSION__)
int variant2122_1(int value, ...) {
    return value + 50;
}
#else
int variant2122_1(int value) {
    return value + 50;
}
#endif
#define MASK2122_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2122_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2122_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2122_3(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 64;
}
int gcd2122_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2122_4(int value) {
    return gcd2122_4(value, 23) * 71;
}
#if defined(__STDC_VERSION__)
int variant2122_5(int value, ...) {
    return value + 78;
}
#else
int variant2122_5(int value) {
    return value + 78;
}
#endif
