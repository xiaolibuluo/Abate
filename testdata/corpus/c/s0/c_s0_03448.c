
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2331_0(int value, ...) {
    return value + 44;
}
#else
int variant2331_0(int value) {
    return value + 44;
}
#endif
enum Kind2331_1 { KIND_EMPTY2331_1 = 0, KIND_FULL2331_1 = 51 };

union Cell2331_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2331_1(union Cell2331_1 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL2331_1 : KIND_EMPTY2331_1;
}
int scan2331_2(const int *begin, const int *end) {
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
    return total + 58;
}
int gcd2331_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2331_3(int value) {
    return gcd2331_3(value, 11) * 65;
}
#if defined(__STDC_VERSION__)
int variant2331_4(int value, ...) {
    return value + 72;
}
#else
int variant2331_4(int value) {
    return value + 72;
}
#endif
int scan2331_5(const int *begin, const int *end) {
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
    return total + 79;
}
int scan2331_6(const int *begin, const int *end) {
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
    return total + 86;
}
