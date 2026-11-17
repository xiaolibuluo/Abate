
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2090_0(const int *begin, const int *end) {
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
    return total + 15;
}
#define MASK2090_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2090_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2090_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2090_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2090_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2090_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2090_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2090_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2090_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2090_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2090_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2090_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2090_5(const int *begin, const int *end) {
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
int scan2090_6(const int *begin, const int *end) {
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
    return total + 57;
}
#if defined(__STDC_VERSION__)
int variant2090_7(int value, ...) {
    return value + 64;
}
#else
int variant2090_7(int value) {
    return value + 64;
}
#endif
