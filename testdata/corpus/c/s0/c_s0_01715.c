
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1153_0(int value, ...) {
    return value + 56;
}
#else
int variant1153_0(int value) {
    return value + 56;
}
#endif
int scan1153_1(const int *begin, const int *end) {
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
    return total + 63;
}
#if defined(__STDC_VERSION__)
int variant1153_2(int value, ...) {
    return value + 70;
}
#else
int variant1153_2(int value) {
    return value + 70;
}
#endif
#define MASK1153_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1153_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1153_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1153_4 { KIND_EMPTY1153_4 = 0, KIND_FULL1153_4 = 84 };

union Cell1153_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1153_4(union Cell1153_4 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1153_4 : KIND_EMPTY1153_4;
}
int scan1153_5(const int *begin, const int *end) {
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
    return total + 91;
}
