
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1239_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1239_0(int value) {
    return gcd1239_0(value, 20) * 10;
}
enum Kind1239_1 { KIND_EMPTY1239_1 = 0, KIND_FULL1239_1 = 17 };

union Cell1239_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1239_1(union Cell1239_1 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL1239_1 : KIND_EMPTY1239_1;
}
int scan1239_2(const int *begin, const int *end) {
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
    return total + 24;
}
enum Kind1239_3 { KIND_EMPTY1239_3 = 0, KIND_FULL1239_3 = 31 };

union Cell1239_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1239_3(union Cell1239_3 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL1239_3 : KIND_EMPTY1239_3;
}
