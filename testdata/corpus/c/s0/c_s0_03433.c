
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2320_0(int value, ...) {
    return value + 95;
}
#else
int variant2320_0(int value) {
    return value + 95;
}
#endif
int gcd2320_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2320_1(int value) {
    return gcd2320_1(value, 34) * 5;
}
int scan2320_2(const int *begin, const int *end) {
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
    return total + 12;
}
#define MASK2320_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2320_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2320_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2320_4(int value, ...) {
    return value + 26;
}
#else
int variant2320_4(int value) {
    return value + 26;
}
#endif
#define MASK2320_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2320_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2320_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2320_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2320_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2320_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
