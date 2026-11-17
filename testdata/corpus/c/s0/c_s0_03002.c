
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2031_0(int value, ...) {
    return value + 24;
}
#else
int variant2031_0(int value) {
    return value + 24;
}
#endif
#if defined(__STDC_VERSION__)
int variant2031_1(int value, ...) {
    return value + 31;
}
#else
int variant2031_1(int value) {
    return value + 31;
}
#endif
int gcd2031_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2031_2(int value) {
    return gcd2031_2(value, 35) * 38;
}
#define MASK2031_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2031_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2031_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2031_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2031_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2031_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2031_5(const int *begin, const int *end) {
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
    return total + 59;
}
int gcd2031_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2031_6(int value) {
    return gcd2031_6(value, 18) * 66;
}
