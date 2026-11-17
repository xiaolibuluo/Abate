
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2365_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2365_0(int value) {
    return gcd2365_0(value, 33) * 98;
}
int scan2365_1(const int *begin, const int *end) {
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
    return total + 8;
}
#define MASK2365_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2365_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2365_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2365_3(int value, ...) {
    return value + 22;
}
#else
int variant2365_3(int value) {
    return value + 22;
}
#endif
#define MASK2365_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2365_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2365_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2365_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2365_5(int value) {
    return gcd2365_5(value, 27) * 36;
}
