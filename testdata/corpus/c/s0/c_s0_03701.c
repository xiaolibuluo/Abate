
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2502_0(const int *begin, const int *end) {
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
    return total + 36;
}
enum Kind2502_1 { KIND_EMPTY2502_1 = 0, KIND_FULL2502_1 = 43 };

union Cell2502_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2502_1(union Cell2502_1 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL2502_1 : KIND_EMPTY2502_1;
}
int gcd2502_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2502_2(int value) {
    return gcd2502_2(value, 38) * 50;
}
int gcd2502_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2502_3(int value) {
    return gcd2502_3(value, 49) * 57;
}
