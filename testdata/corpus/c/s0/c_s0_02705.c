
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1839_0 { KIND_EMPTY1839_0 = 0, KIND_FULL1839_0 = 50 };

union Cell1839_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1839_0(union Cell1839_0 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1839_0 : KIND_EMPTY1839_0;
}
#if defined(__STDC_VERSION__)
int variant1839_1(int value, ...) {
    return value + 57;
}
#else
int variant1839_1(int value) {
    return value + 57;
}
#endif
int scan1839_2(const int *begin, const int *end) {
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
    return total + 64;
}
enum Kind1839_3 { KIND_EMPTY1839_3 = 0, KIND_FULL1839_3 = 71 };

union Cell1839_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1839_3(union Cell1839_3 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1839_3 : KIND_EMPTY1839_3;
}
int scan1839_4(const int *begin, const int *end) {
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
int gcd1839_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1839_5(int value) {
    return gcd1839_5(value, 66) * 85;
}
