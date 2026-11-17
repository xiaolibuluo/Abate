
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2556_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2556_0(int value) {
    return gcd2556_0(value, 28) * 59;
}
enum Kind2556_1 { KIND_EMPTY2556_1 = 0, KIND_FULL2556_1 = 66 };

union Cell2556_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2556_1(union Cell2556_1 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL2556_1 : KIND_EMPTY2556_1;
}
enum Kind2556_2 { KIND_EMPTY2556_2 = 0, KIND_FULL2556_2 = 73 };

union Cell2556_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2556_2(union Cell2556_2 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL2556_2 : KIND_EMPTY2556_2;
}
int scan2556_3(const int *begin, const int *end) {
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
    return total + 80;
}
