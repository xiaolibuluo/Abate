
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant109_0(int value, ...) {
    return value + 64;
}
#else
int variant109_0(int value) {
    return value + 64;
}
#endif
#if defined(__STDC_VERSION__)
int variant109_1(int value, ...) {
    return value + 71;
}
#else
int variant109_1(int value) {
    return value + 71;
}
#endif
int scan109_2(const int *begin, const int *end) {
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
    return total + 78;
}
