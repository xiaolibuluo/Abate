
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2597_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2597_0(int value) {
    return gcd2597_0(value, 10) * 10;
}
int scan2597_1(const int *begin, const int *end) {
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
    return total + 17;
}
#if defined(__STDC_VERSION__)
int variant2597_2(int value, ...) {
    return value + 24;
}
#else
int variant2597_2(int value) {
    return value + 24;
}
#endif
#define MASK2597_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2597_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2597_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
