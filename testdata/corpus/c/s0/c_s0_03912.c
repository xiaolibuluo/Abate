
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2650_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2650_0(int value) {
    return gcd2650_0(value, 15) * 20;
}
#if defined(__STDC_VERSION__)
int variant2650_1(int value, ...) {
    return value + 27;
}
#else
int variant2650_1(int value) {
    return value + 27;
}
#endif
#if defined(__STDC_VERSION__)
int variant2650_2(int value, ...) {
    return value + 34;
}
#else
int variant2650_2(int value) {
    return value + 34;
}
#endif
int scan2650_3(const int *begin, const int *end) {
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
    return total + 41;
}
#if defined(__STDC_VERSION__)
int variant2650_4(int value, ...) {
    return value + 48;
}
#else
int variant2650_4(int value) {
    return value + 48;
}
#endif
