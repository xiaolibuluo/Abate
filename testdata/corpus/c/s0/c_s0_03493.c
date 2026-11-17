
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2357_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2357_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2357_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2357_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2357_1(int value) {
    return gcd2357_1(value, 49) * 98;
}
#define MASK2357_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2357_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2357_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2357_3(int value, ...) {
    return value + 15;
}
#else
int variant2357_3(int value) {
    return value + 15;
}
#endif
int scan2357_4(const int *begin, const int *end) {
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
    return total + 22;
}
