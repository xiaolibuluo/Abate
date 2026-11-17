
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1142_0(int value, ...) {
    return value + 10;
}
#else
int variant1142_0(int value) {
    return value + 10;
}
#endif
int scan1142_1(const int *begin, const int *end) {
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
    return total + 17;
}
enum Kind1142_2 { KIND_EMPTY1142_2 = 0, KIND_FULL1142_2 = 24 };

union Cell1142_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1142_2(union Cell1142_2 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1142_2 : KIND_EMPTY1142_2;
}
int gcd1142_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1142_3(int value) {
    return gcd1142_3(value, 45) * 31;
}
