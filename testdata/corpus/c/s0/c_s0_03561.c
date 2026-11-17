
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2403_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2403_0(int value) {
    return gcd2403_0(value, 55) * 10;
}
int scan2403_1(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant2403_2(int value, ...) {
    return value + 24;
}
#else
int variant2403_2(int value) {
    return value + 24;
}
#endif
enum Kind2403_3 { KIND_EMPTY2403_3 = 0, KIND_FULL2403_3 = 31 };

union Cell2403_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2403_3(union Cell2403_3 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL2403_3 : KIND_EMPTY2403_3;
}
