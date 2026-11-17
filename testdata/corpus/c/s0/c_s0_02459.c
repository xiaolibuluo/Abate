
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1670_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1670_0(int value) {
    return gcd1670_0(value, 48) * 84;
}
enum Kind1670_1 { KIND_EMPTY1670_1 = 0, KIND_FULL1670_1 = 91 };

union Cell1670_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1670_1(union Cell1670_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1670_1 : KIND_EMPTY1670_1;
}
int scan1670_2(const int *begin, const int *end) {
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
