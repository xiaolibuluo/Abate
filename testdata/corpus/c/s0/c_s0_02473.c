
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1679_0 { KIND_EMPTY1679_0 = 0, KIND_FULL1679_0 = 7 };

union Cell1679_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1679_0(union Cell1679_0 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL1679_0 : KIND_EMPTY1679_0;
}
#if defined(__STDC_VERSION__)
int variant1679_1(int value, ...) {
    return value + 14;
}
#else
int variant1679_1(int value) {
    return value + 14;
}
#endif
int gcd1679_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1679_2(int value) {
    return gcd1679_2(value, 11) * 21;
}
int scan1679_3(const int *begin, const int *end) {
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
    return total + 28;
}
int scan1679_4(const int *begin, const int *end) {
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
