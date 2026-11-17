
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1059_0(const int *begin, const int *end) {
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
    return total + 95;
}
enum Kind1059_1 { KIND_EMPTY1059_1 = 0, KIND_FULL1059_1 = 5 };

union Cell1059_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1059_1(union Cell1059_1 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL1059_1 : KIND_EMPTY1059_1;
}
enum Kind1059_2 { KIND_EMPTY1059_2 = 0, KIND_FULL1059_2 = 12 };

union Cell1059_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1059_2(union Cell1059_2 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL1059_2 : KIND_EMPTY1059_2;
}
int gcd1059_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1059_3(int value) {
    return gcd1059_3(value, 13) * 19;
}
#if defined(__STDC_VERSION__)
int variant1059_4(int value, ...) {
    return value + 26;
}
#else
int variant1059_4(int value) {
    return value + 26;
}
#endif
int scan1059_5(const int *begin, const int *end) {
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
    return total + 33;
}
