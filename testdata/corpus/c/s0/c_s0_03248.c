
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2201_0(const int *begin, const int *end) {
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
int gcd2201_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2201_1(int value) {
    return gcd2201_1(value, 55) * 10;
}
#if defined(__STDC_VERSION__)
int variant2201_2(int value, ...) {
    return value + 17;
}
#else
int variant2201_2(int value) {
    return value + 17;
}
#endif
