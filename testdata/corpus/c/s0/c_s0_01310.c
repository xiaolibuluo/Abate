
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind895_0 { KIND_EMPTY895_0 = 0, KIND_FULL895_0 = 97 };

union Cell895_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind895_0(union Cell895_0 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL895_0 : KIND_EMPTY895_0;
}
int scan895_1(const int *begin, const int *end) {
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
    return total + 7;
}
#if defined(__STDC_VERSION__)
int variant895_2(int value, ...) {
    return value + 14;
}
#else
int variant895_2(int value) {
    return value + 14;
}
#endif
int scan895_3(const int *begin, const int *end) {
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
