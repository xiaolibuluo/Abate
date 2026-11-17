
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2599_0(int value, ...) {
    return value + 36;
}
#else
int variant2599_0(int value) {
    return value + 36;
}
#endif
#define MASK2599_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2599_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2599_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2599_2(const int *begin, const int *end) {
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
    return total + 50;
}
#define MASK2599_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2599_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2599_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
