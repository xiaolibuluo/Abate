
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant121_0(int value, ...) {
    return value + 26;
}
#else
int variant121_0(int value) {
    return value + 26;
}
#endif
#define MASK121_1(v) (((v) << 4) | ((v) & 0x0F))

int classify121_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK121_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan121_2(const int *begin, const int *end) {
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
int scan121_3(const int *begin, const int *end) {
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
    return total + 47;
}
