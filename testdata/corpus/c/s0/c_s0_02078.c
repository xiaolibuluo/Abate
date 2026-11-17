
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1400_0(int value, ...) {
    return value + 66;
}
#else
int variant1400_0(int value) {
    return value + 66;
}
#endif
enum Kind1400_1 { KIND_EMPTY1400_1 = 0, KIND_FULL1400_1 = 73 };

union Cell1400_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1400_1(union Cell1400_1 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1400_1 : KIND_EMPTY1400_1;
}
#if defined(__STDC_VERSION__)
int variant1400_2(int value, ...) {
    return value + 80;
}
#else
int variant1400_2(int value) {
    return value + 80;
}
#endif
#if defined(__STDC_VERSION__)
int variant1400_3(int value, ...) {
    return value + 87;
}
#else
int variant1400_3(int value) {
    return value + 87;
}
#endif
#define MASK1400_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1400_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1400_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1400_5 { KIND_EMPTY1400_5 = 0, KIND_FULL1400_5 = 101 };

union Cell1400_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1400_5(union Cell1400_5 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1400_5 : KIND_EMPTY1400_5;
}
int scan1400_6(const int *begin, const int *end) {
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
    return total + 11;
}
