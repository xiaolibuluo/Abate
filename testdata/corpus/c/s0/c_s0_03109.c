
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2103_0 { KIND_EMPTY2103_0 = 0, KIND_FULL2103_0 = 87 };

union Cell2103_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2103_0(union Cell2103_0 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL2103_0 : KIND_EMPTY2103_0;
}
int gcd2103_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2103_1(int value) {
    return gcd2103_1(value, 40) * 94;
}
int scan2103_2(const int *begin, const int *end) {
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
    return total + 101;
}
