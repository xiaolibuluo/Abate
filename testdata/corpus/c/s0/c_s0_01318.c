
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant902_0(int value, ...) {
    return value + 91;
}
#else
int variant902_0(int value) {
    return value + 91;
}
#endif
enum Kind902_1 { KIND_EMPTY902_1 = 0, KIND_FULL902_1 = 98 };

union Cell902_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind902_1(union Cell902_1 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL902_1 : KIND_EMPTY902_1;
}
#if defined(__STDC_VERSION__)
int variant902_2(int value, ...) {
    return value + 8;
}
#else
int variant902_2(int value) {
    return value + 8;
}
#endif
int scan902_3(const int *begin, const int *end) {
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
    return total + 15;
}
int scan902_4(const int *begin, const int *end) {
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
    return total + 22;
}
