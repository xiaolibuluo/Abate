
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1725_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1725_0(int value) {
    return gcd1725_0(value, 67) * 23;
}
enum Kind1725_1 { KIND_EMPTY1725_1 = 0, KIND_FULL1725_1 = 30 };

union Cell1725_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1725_1(union Cell1725_1 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1725_1 : KIND_EMPTY1725_1;
}
#if defined(__STDC_VERSION__)
int variant1725_2(int value, ...) {
    return value + 37;
}
#else
int variant1725_2(int value) {
    return value + 37;
}
#endif
#if defined(__STDC_VERSION__)
int variant1725_3(int value, ...) {
    return value + 44;
}
#else
int variant1725_3(int value) {
    return value + 44;
}
#endif
int scan1725_4(const int *begin, const int *end) {
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
    return total + 51;
}
