
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1438_0(const int *begin, const int *end) {
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
    return total + 75;
}
#if defined(__STDC_VERSION__)
int variant1438_1(int value, ...) {
    return value + 82;
}
#else
int variant1438_1(int value) {
    return value + 82;
}
#endif
enum Kind1438_2 { KIND_EMPTY1438_2 = 0, KIND_FULL1438_2 = 89 };

union Cell1438_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1438_2(union Cell1438_2 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1438_2 : KIND_EMPTY1438_2;
}
int scan1438_3(const int *begin, const int *end) {
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
#define MASK1438_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1438_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1438_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
