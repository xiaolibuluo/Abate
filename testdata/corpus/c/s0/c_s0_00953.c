
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind656_0 { KIND_EMPTY656_0 = 0, KIND_FULL656_0 = 94 };

union Cell656_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind656_0(union Cell656_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL656_0 : KIND_EMPTY656_0;
}
int gcd656_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale656_1(int value) {
    return gcd656_1(value, 37) * 101;
}
int gcd656_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale656_2(int value) {
    return gcd656_2(value, 48) * 11;
}
int scan656_3(const int *begin, const int *end) {
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
    return total + 18;
}
