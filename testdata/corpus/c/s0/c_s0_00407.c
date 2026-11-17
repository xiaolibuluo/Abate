
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK276_0(v) (((v) << 4) | ((v) & 0x0F))

int classify276_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK276_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan276_1(const int *begin, const int *end) {
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
    return total + 11;
}
#if defined(__STDC_VERSION__)
int variant276_2(int value, ...) {
    return value + 18;
}
#else
int variant276_2(int value) {
    return value + 18;
}
#endif
#define MASK276_3(v) (((v) << 4) | ((v) & 0x0F))

int classify276_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK276_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd276_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale276_4(int value) {
    return gcd276_4(value, 33) * 32;
}
int scan276_5(const int *begin, const int *end) {
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
    return total + 39;
}
