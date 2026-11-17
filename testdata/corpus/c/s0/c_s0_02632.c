
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1786_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1786_0(int value) {
    return gcd1786_0(value, 67) * 40;
}
enum Kind1786_1 { KIND_EMPTY1786_1 = 0, KIND_FULL1786_1 = 47 };

union Cell1786_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1786_1(union Cell1786_1 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1786_1 : KIND_EMPTY1786_1;
}
#if defined(__STDC_VERSION__)
int variant1786_2(int value, ...) {
    return value + 54;
}
#else
int variant1786_2(int value) {
    return value + 54;
}
#endif
enum Kind1786_3 { KIND_EMPTY1786_3 = 0, KIND_FULL1786_3 = 61 };

union Cell1786_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1786_3(union Cell1786_3 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1786_3 : KIND_EMPTY1786_3;
}
#if defined(__STDC_VERSION__)
int variant1786_4(int value, ...) {
    return value + 68;
}
#else
int variant1786_4(int value) {
    return value + 68;
}
#endif
int scan1786_5(const int *begin, const int *end) {
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
    return total + 75;
}
