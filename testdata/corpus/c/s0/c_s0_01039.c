
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind712_0 { KIND_EMPTY712_0 = 0, KIND_FULL712_0 = 46 };

union Cell712_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind712_0(union Cell712_0 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL712_0 : KIND_EMPTY712_0;
}
int gcd712_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale712_1(int value) {
    return gcd712_1(value, 63) * 53;
}
int gcd712_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale712_2(int value) {
    return gcd712_2(value, 13) * 60;
}
#if defined(__STDC_VERSION__)
int variant712_3(int value, ...) {
    return value + 67;
}
#else
int variant712_3(int value) {
    return value + 67;
}
#endif
int scan712_4(const int *begin, const int *end) {
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
    return total + 74;
}
int gcd712_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale712_5(int value) {
    return gcd712_5(value, 46) * 81;
}
