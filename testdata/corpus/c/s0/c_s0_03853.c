
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2604_0 { KIND_EMPTY2604_0 = 0, KIND_FULL2604_0 = 101 };

union Cell2604_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2604_0(union Cell2604_0 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL2604_0 : KIND_EMPTY2604_0;
}
int scan2604_1(const int *begin, const int *end) {
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
    return total + 11;
}
int gcd2604_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2604_2(int value) {
    return gcd2604_2(value, 20) * 18;
}
int scan2604_3(const int *begin, const int *end) {
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
    return total + 25;
}
