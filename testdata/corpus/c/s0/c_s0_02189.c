
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1473_0(const int *begin, const int *end) {
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
    return total + 45;
}
int gcd1473_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1473_1(int value) {
    return gcd1473_1(value, 22) * 52;
}
enum Kind1473_2 { KIND_EMPTY1473_2 = 0, KIND_FULL1473_2 = 59 };

union Cell1473_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1473_2(union Cell1473_2 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL1473_2 : KIND_EMPTY1473_2;
}
