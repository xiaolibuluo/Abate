
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1299_0 { KIND_EMPTY1299_0 = 0, KIND_FULL1299_0 = 14 };

union Cell1299_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1299_0(union Cell1299_0 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL1299_0 : KIND_EMPTY1299_0;
}
int gcd1299_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1299_1(int value) {
    return gcd1299_1(value, 24) * 21;
}
int gcd1299_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1299_2(int value) {
    return gcd1299_2(value, 35) * 28;
}
int scan1299_3(const int *begin, const int *end) {
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
    return total + 35;
}
#if defined(__STDC_VERSION__)
int variant1299_4(int value, ...) {
    return value + 42;
}
#else
int variant1299_4(int value) {
    return value + 42;
}
#endif
int gcd1299_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1299_5(int value) {
    return gcd1299_5(value, 68) * 49;
}
