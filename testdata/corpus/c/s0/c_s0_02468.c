
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1674_0(int value, ...) {
    return value + 39;
}
#else
int variant1674_0(int value) {
    return value + 39;
}
#endif
int gcd1674_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1674_1(int value) {
    return gcd1674_1(value, 26) * 46;
}
#if defined(__STDC_VERSION__)
int variant1674_2(int value, ...) {
    return value + 53;
}
#else
int variant1674_2(int value) {
    return value + 53;
}
#endif
int scan1674_3(const int *begin, const int *end) {
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
    return total + 60;
}
