
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1842_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1842_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1842_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1842_1(const int *begin, const int *end) {
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
    return total + 96;
}
#if defined(__STDC_VERSION__)
int variant1842_2(int value, ...) {
    return value + 6;
}
#else
int variant1842_2(int value) {
    return value + 6;
}
#endif
int scan1842_3(const int *begin, const int *end) {
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
    return total + 13;
}
#define MASK1842_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1842_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1842_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1842_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1842_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1842_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1842_6 { KIND_EMPTY1842_6 = 0, KIND_FULL1842_6 = 34 };

union Cell1842_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1842_6(union Cell1842_6 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1842_6 : KIND_EMPTY1842_6;
}
