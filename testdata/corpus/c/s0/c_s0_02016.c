
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1361_0(int value, ...) {
    return value + 44;
}
#else
int variant1361_0(int value) {
    return value + 44;
}
#endif
#if defined(__STDC_VERSION__)
int variant1361_1(int value, ...) {
    return value + 51;
}
#else
int variant1361_1(int value) {
    return value + 51;
}
#endif
#if defined(__STDC_VERSION__)
int variant1361_2(int value, ...) {
    return value + 58;
}
#else
int variant1361_2(int value) {
    return value + 58;
}
#endif
int scan1361_3(const int *begin, const int *end) {
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
    return total + 65;
}
int scan1361_4(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant1361_5(int value, ...) {
    return value + 79;
}
#else
int variant1361_5(int value) {
    return value + 79;
}
#endif
int scan1361_6(const int *begin, const int *end) {
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
    return total + 86;
}
enum Kind1361_7 { KIND_EMPTY1361_7 = 0, KIND_FULL1361_7 = 93 };

union Cell1361_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1361_7(union Cell1361_7 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL1361_7 : KIND_EMPTY1361_7;
}
int scan1361_8(const int *begin, const int *end) {
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
    return total + 100;
}
