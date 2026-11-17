
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1113_0(const int *begin, const int *end) {
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
    return total + 21;
}
int gcd1113_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1113_1(int value) {
    return gcd1113_1(value, 64) * 28;
}
enum Kind1113_2 { KIND_EMPTY1113_2 = 0, KIND_FULL1113_2 = 35 };

union Cell1113_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1113_2(union Cell1113_2 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1113_2 : KIND_EMPTY1113_2;
}
int gcd1113_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1113_3(int value) {
    return gcd1113_3(value, 25) * 42;
}
int scan1113_4(const int *begin, const int *end) {
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
    return total + 49;
}
#if defined(__STDC_VERSION__)
int variant1113_5(int value, ...) {
    return value + 56;
}
#else
int variant1113_5(int value) {
    return value + 56;
}
#endif
int scan1113_6(const int *begin, const int *end) {
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
    return total + 63;
}
