
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2455_0(int value, ...) {
    return value + 7;
}
#else
int variant2455_0(int value) {
    return value + 7;
}
#endif
#define MASK2455_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2455_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2455_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2455_2(const int *begin, const int *end) {
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
    return total + 21;
}
int scan2455_3(const int *begin, const int *end) {
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
    return total + 28;
}
