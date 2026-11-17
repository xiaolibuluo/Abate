
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1844_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1844_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1844_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1844_1(const int *begin, const int *end) {
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
    return total + 25;
}
int gcd1844_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1844_2(int value) {
    return gcd1844_2(value, 68) * 32;
}
#if defined(__STDC_VERSION__)
int variant1844_3(int value, ...) {
    return value + 39;
}
#else
int variant1844_3(int value) {
    return value + 39;
}
#endif
int scan1844_4(const int *begin, const int *end) {
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
    return total + 46;
}
int gcd1844_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1844_5(int value) {
    return gcd1844_5(value, 40) * 53;
}
int scan1844_6(const int *begin, const int *end) {
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
    return total + 60;
}
