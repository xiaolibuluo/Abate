
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1832_0(const int *begin, const int *end) {
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
    return total + 56;
}
#if defined(__STDC_VERSION__)
int variant1832_1(int value, ...) {
    return value + 63;
}
#else
int variant1832_1(int value) {
    return value + 63;
}
#endif
#if defined(__STDC_VERSION__)
int variant1832_2(int value, ...) {
    return value + 70;
}
#else
int variant1832_2(int value) {
    return value + 70;
}
#endif
int scan1832_3(const int *begin, const int *end) {
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
    return total + 77;
}
int scan1832_4(const int *begin, const int *end) {
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
    return total + 84;
}
#if defined(__STDC_VERSION__)
int variant1832_5(int value, ...) {
    return value + 91;
}
#else
int variant1832_5(int value) {
    return value + 91;
}
#endif
