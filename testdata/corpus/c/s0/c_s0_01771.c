
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1188_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1188_0(int value) {
    return gcd1188_0(value, 29) * 26;
}
#define MASK1188_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1188_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1188_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1188_2(const int *begin, const int *end) {
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
    return total + 40;
}
#if defined(__STDC_VERSION__)
int variant1188_3(int value, ...) {
    return value + 47;
}
#else
int variant1188_3(int value) {
    return value + 47;
}
#endif
