
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind602_0 { KIND_EMPTY602_0 = 0, KIND_FULL602_0 = 71 };

union Cell602_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind602_0(union Cell602_0 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL602_0 : KIND_EMPTY602_0;
}
int scan602_1(const int *begin, const int *end) {
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
    return total + 78;
}
int gcd602_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale602_2(int value) {
    return gcd602_2(value, 36) * 85;
}
