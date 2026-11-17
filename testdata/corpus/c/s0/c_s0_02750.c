
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1872_0 { KIND_EMPTY1872_0 = 0, KIND_FULL1872_0 = 91 };

union Cell1872_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1872_0(union Cell1872_0 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1872_0 : KIND_EMPTY1872_0;
}
int scan1872_1(const int *begin, const int *end) {
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
    return total + 98;
}
int scan1872_2(const int *begin, const int *end) {
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
    return total + 8;
}
int gcd1872_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1872_3(int value) {
    return gcd1872_3(value, 31) * 15;
}
