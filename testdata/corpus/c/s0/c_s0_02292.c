
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1548_0(const int *begin, const int *end) {
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
enum Kind1548_1 { KIND_EMPTY1548_1 = 0, KIND_FULL1548_1 = 57 };

union Cell1548_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1548_1(union Cell1548_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1548_1 : KIND_EMPTY1548_1;
}
#define MASK1548_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1548_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1548_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1548_3(int value, ...) {
    return value + 71;
}
#else
int variant1548_3(int value) {
    return value + 71;
}
#endif
#define MASK1548_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1548_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1548_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1548_5(const int *begin, const int *end) {
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
    return total + 85;
}
#define MASK1548_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1548_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1548_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
