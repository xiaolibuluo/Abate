
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1353_0(const int *begin, const int *end) {
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
    return total + 37;
}
#define MASK1353_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1353_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1353_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1353_2(int value, ...) {
    return value + 51;
}
#else
int variant1353_2(int value) {
    return value + 51;
}
#endif
int scan1353_3(const int *begin, const int *end) {
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
    return total + 58;
}
int scan1353_4(const int *begin, const int *end) {
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
    return total + 65;
}
int gcd1353_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1353_5(int value) {
    return gcd1353_5(value, 19) * 72;
}
#define MASK1353_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1353_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1353_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
