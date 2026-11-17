
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2158_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2158_0(int value) {
    return gcd2158_0(value, 48) * 26;
}
enum Kind2158_1 { KIND_EMPTY2158_1 = 0, KIND_FULL2158_1 = 33 };

union Cell2158_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2158_1(union Cell2158_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL2158_1 : KIND_EMPTY2158_1;
}
int scan2158_2(const int *begin, const int *end) {
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
    return total + 40;
}
