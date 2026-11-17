
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan600_0(const int *begin, const int *end) {
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
int gcd600_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale600_1(int value) {
    return gcd600_1(value, 11) * 52;
}
enum Kind600_2 { KIND_EMPTY600_2 = 0, KIND_FULL600_2 = 59 };

union Cell600_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind600_2(union Cell600_2 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL600_2 : KIND_EMPTY600_2;
}
