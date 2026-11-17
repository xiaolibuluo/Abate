
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1426_0 { KIND_EMPTY1426_0 = 0, KIND_FULL1426_0 = 16 };

union Cell1426_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1426_0(union Cell1426_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1426_0 : KIND_EMPTY1426_0;
}
#if defined(__STDC_VERSION__)
int variant1426_1(int value, ...) {
    return value + 23;
}
#else
int variant1426_1(int value) {
    return value + 23;
}
#endif
int scan1426_2(const int *begin, const int *end) {
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
    return total + 30;
}
