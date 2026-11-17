
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1962_0(const int *begin, const int *end) {
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
    return total + 97;
}
enum Kind1962_1 { KIND_EMPTY1962_1 = 0, KIND_FULL1962_1 = 7 };

union Cell1962_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1962_1(union Cell1962_1 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1962_1 : KIND_EMPTY1962_1;
}
int gcd1962_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1962_2(int value) {
    return gcd1962_2(value, 40) * 14;
}
