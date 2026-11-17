
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1377_0(int value, ...) {
    return value + 58;
}
#else
int variant1377_0(int value) {
    return value + 58;
}
#endif
#if defined(__STDC_VERSION__)
int variant1377_1(int value, ...) {
    return value + 65;
}
#else
int variant1377_1(int value) {
    return value + 65;
}
#endif
int scan1377_2(const int *begin, const int *end) {
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
    return total + 72;
}
#define MASK1377_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1377_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1377_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
