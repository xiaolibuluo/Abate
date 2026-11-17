
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1586_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1586_0(int value) {
    return gcd1586_0(value, 9) * 59;
}
int gcd1586_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1586_1(int value) {
    return gcd1586_1(value, 20) * 66;
}
enum Kind1586_2 { KIND_EMPTY1586_2 = 0, KIND_FULL1586_2 = 73 };

union Cell1586_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1586_2(union Cell1586_2 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL1586_2 : KIND_EMPTY1586_2;
}
int scan1586_3(const int *begin, const int *end) {
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
int gcd1586_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1586_4(int value) {
    return gcd1586_4(value, 53) * 87;
}
