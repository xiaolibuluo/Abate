
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan574_0(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant574_1(int value, ...) {
    return value + 5;
}
#else
int variant574_1(int value) {
    return value + 5;
}
#endif
enum Kind574_2 { KIND_EMPTY574_2 = 0, KIND_FULL574_2 = 12 };

union Cell574_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind574_2(union Cell574_2 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL574_2 : KIND_EMPTY574_2;
}
int gcd574_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale574_3(int value) {
    return gcd574_3(value, 34) * 19;
}
int scan574_4(const int *begin, const int *end) {
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
    return total + 26;
}
int scan574_5(const int *begin, const int *end) {
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
