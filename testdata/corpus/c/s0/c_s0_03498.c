
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2359_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2359_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2359_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2359_1(const int *begin, const int *end) {
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
    return total + 27;
}
#if defined(__STDC_VERSION__)
int variant2359_2(int value, ...) {
    return value + 34;
}
#else
int variant2359_2(int value) {
    return value + 34;
}
#endif
int gcd2359_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2359_3(int value) {
    return gcd2359_3(value, 24) * 41;
}
#define MASK2359_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2359_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2359_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
