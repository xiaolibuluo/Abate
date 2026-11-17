
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1629_0(int value, ...) {
    return value + 36;
}
#else
int variant1629_0(int value) {
    return value + 36;
}
#endif
int gcd1629_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1629_1(int value) {
    return gcd1629_1(value, 16) * 43;
}
enum Kind1629_2 { KIND_EMPTY1629_2 = 0, KIND_FULL1629_2 = 50 };

union Cell1629_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1629_2(union Cell1629_2 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL1629_2 : KIND_EMPTY1629_2;
}
#if defined(__STDC_VERSION__)
int variant1629_3(int value, ...) {
    return value + 57;
}
#else
int variant1629_3(int value) {
    return value + 57;
}
#endif
int scan1629_4(const int *begin, const int *end) {
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
    return total + 64;
}
