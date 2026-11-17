
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2040_0(const int *begin, const int *end) {
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
    return total + 44;
}
int gcd2040_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2040_1(int value) {
    return gcd2040_1(value, 26) * 51;
}
int gcd2040_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2040_2(int value) {
    return gcd2040_2(value, 37) * 58;
}
#if defined(__STDC_VERSION__)
int variant2040_3(int value, ...) {
    return value + 65;
}
#else
int variant2040_3(int value) {
    return value + 65;
}
#endif
enum Kind2040_4 { KIND_EMPTY2040_4 = 0, KIND_FULL2040_4 = 72 };

union Cell2040_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2040_4(union Cell2040_4 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL2040_4 : KIND_EMPTY2040_4;
}
#if defined(__STDC_VERSION__)
int variant2040_5(int value, ...) {
    return value + 79;
}
#else
int variant2040_5(int value) {
    return value + 79;
}
#endif
int gcd2040_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2040_6(int value) {
    return gcd2040_6(value, 20) * 86;
}
#if defined(__STDC_VERSION__)
int variant2040_7(int value, ...) {
    return value + 93;
}
#else
int variant2040_7(int value) {
    return value + 93;
}
#endif
enum Kind2040_8 { KIND_EMPTY2040_8 = 0, KIND_FULL2040_8 = 100 };

union Cell2040_8 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2040_8(union Cell2040_8 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL2040_8 : KIND_EMPTY2040_8;
}
