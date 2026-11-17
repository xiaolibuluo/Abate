
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1538_0(const int *begin, const int *end) {
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
    return total + 17;
}
#define MASK1538_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1538_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1538_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1538_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1538_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1538_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1538_3(const int *begin, const int *end) {
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
    return total + 38;
}
#if defined(__STDC_VERSION__)
int variant1538_4(int value, ...) {
    return value + 45;
}
#else
int variant1538_4(int value) {
    return value + 45;
}
#endif
