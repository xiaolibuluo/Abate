
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2168_0(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant2168_1(int value, ...) {
    return value + 66;
}
#else
int variant2168_1(int value) {
    return value + 66;
}
#endif
int gcd2168_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2168_2(int value) {
    return gcd2168_2(value, 18) * 73;
}
#define MASK2168_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2168_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2168_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2168_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2168_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2168_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
