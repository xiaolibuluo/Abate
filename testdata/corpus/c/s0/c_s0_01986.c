
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1341_0(const int *begin, const int *end) {
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
    return total + 75;
}
#if defined(__STDC_VERSION__)
int variant1341_1(int value, ...) {
    return value + 82;
}
#else
int variant1341_1(int value) {
    return value + 82;
}
#endif
#define MASK1341_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1341_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1341_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1341_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1341_3(int value) {
    return gcd1341_3(value, 35) * 96;
}
int scan1341_4(const int *begin, const int *end) {
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
    return total + 6;
}
