
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd113_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale113_0(int value) {
    return gcd113_0(value, 68) * 19;
}
enum Kind113_1 { KIND_EMPTY113_1 = 0, KIND_FULL113_1 = 26 };

union Cell113_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind113_1(union Cell113_1 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL113_1 : KIND_EMPTY113_1;
}
int scan113_2(const int *begin, const int *end) {
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
    return total + 33;
}
