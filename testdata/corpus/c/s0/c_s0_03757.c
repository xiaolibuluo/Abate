
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2541_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2541_0(int value) {
    return gcd2541_0(value, 45) * 58;
}
#if defined(__STDC_VERSION__)
int variant2541_1(int value, ...) {
    return value + 65;
}
#else
int variant2541_1(int value) {
    return value + 65;
}
#endif
int scan2541_2(const int *begin, const int *end) {
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
