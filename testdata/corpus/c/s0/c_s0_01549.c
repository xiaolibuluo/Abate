
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1047_0 { KIND_EMPTY1047_0 = 0, KIND_FULL1047_0 = 36 };

union Cell1047_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1047_0(union Cell1047_0 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1047_0 : KIND_EMPTY1047_0;
}
int scan1047_1(const int *begin, const int *end) {
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
    return total + 43;
}
int gcd1047_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1047_2(int value) {
    return gcd1047_2(value, 40) * 50;
}
#if defined(__STDC_VERSION__)
int variant1047_3(int value, ...) {
    return value + 57;
}
#else
int variant1047_3(int value) {
    return value + 57;
}
#endif
#if defined(__STDC_VERSION__)
int variant1047_4(int value, ...) {
    return value + 64;
}
#else
int variant1047_4(int value) {
    return value + 64;
}
#endif
