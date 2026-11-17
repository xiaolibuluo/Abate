
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1603_0(const int *begin, const int *end) {
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
int scan1603_1(const int *begin, const int *end) {
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
    return total + 93;
}
enum Kind1603_2 { KIND_EMPTY1603_2 = 0, KIND_FULL1603_2 = 100 };

union Cell1603_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1603_2(union Cell1603_2 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL1603_2 : KIND_EMPTY1603_2;
}
#if defined(__STDC_VERSION__)
int variant1603_3(int value, ...) {
    return value + 10;
}
#else
int variant1603_3(int value) {
    return value + 10;
}
#endif
enum Kind1603_4 { KIND_EMPTY1603_4 = 0, KIND_FULL1603_4 = 17 };

union Cell1603_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1603_4(union Cell1603_4 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL1603_4 : KIND_EMPTY1603_4;
}
enum Kind1603_5 { KIND_EMPTY1603_5 = 0, KIND_FULL1603_5 = 24 };

union Cell1603_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1603_5(union Cell1603_5 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL1603_5 : KIND_EMPTY1603_5;
}
int gcd1603_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1603_6(int value) {
    return gcd1603_6(value, 11) * 31;
}
