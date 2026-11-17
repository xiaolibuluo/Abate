
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1235_0(int value, ...) {
    return value + 55;
}
#else
int variant1235_0(int value) {
    return value + 55;
}
#endif
enum Kind1235_1 { KIND_EMPTY1235_1 = 0, KIND_FULL1235_1 = 62 };

union Cell1235_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1235_1(union Cell1235_1 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1235_1 : KIND_EMPTY1235_1;
}
int scan1235_2(const int *begin, const int *end) {
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
    return total + 69;
}
int gcd1235_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1235_3(int value) {
    return gcd1235_3(value, 25) * 76;
}
