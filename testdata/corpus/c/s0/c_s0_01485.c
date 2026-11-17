
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1008_0 { KIND_EMPTY1008_0 = 0, KIND_FULL1008_0 = 14 };

union Cell1008_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1008_0(union Cell1008_0 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL1008_0 : KIND_EMPTY1008_0;
}
int scan1008_1(const int *begin, const int *end) {
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
    return total + 21;
}
#if defined(__STDC_VERSION__)
int variant1008_2(int value, ...) {
    return value + 28;
}
#else
int variant1008_2(int value) {
    return value + 28;
}
#endif
enum Kind1008_3 { KIND_EMPTY1008_3 = 0, KIND_FULL1008_3 = 35 };

union Cell1008_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1008_3(union Cell1008_3 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1008_3 : KIND_EMPTY1008_3;
}
#if defined(__STDC_VERSION__)
int variant1008_4(int value, ...) {
    return value + 42;
}
#else
int variant1008_4(int value) {
    return value + 42;
}
#endif
enum Kind1008_5 { KIND_EMPTY1008_5 = 0, KIND_FULL1008_5 = 49 };

union Cell1008_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1008_5(union Cell1008_5 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1008_5 : KIND_EMPTY1008_5;
}
