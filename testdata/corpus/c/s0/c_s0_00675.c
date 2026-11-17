
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan463_0(const int *begin, const int *end) {
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
    return total + 10;
}
int gcd463_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale463_1(int value) {
    return gcd463_1(value, 28) * 17;
}
enum Kind463_2 { KIND_EMPTY463_2 = 0, KIND_FULL463_2 = 24 };

union Cell463_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind463_2(union Cell463_2 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL463_2 : KIND_EMPTY463_2;
}
int gcd463_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale463_3(int value) {
    return gcd463_3(value, 50) * 31;
}
