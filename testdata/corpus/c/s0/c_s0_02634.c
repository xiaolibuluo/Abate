
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1787_0(const int *begin, const int *end) {
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
    return total + 53;
}
int scan1787_1(const int *begin, const int *end) {
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
    return total + 60;
}
int gcd1787_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1787_2(int value) {
    return gcd1787_2(value, 35) * 67;
}
enum Kind1787_3 { KIND_EMPTY1787_3 = 0, KIND_FULL1787_3 = 74 };

union Cell1787_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1787_3(union Cell1787_3 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1787_3 : KIND_EMPTY1787_3;
}
enum Kind1787_4 { KIND_EMPTY1787_4 = 0, KIND_FULL1787_4 = 81 };

union Cell1787_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1787_4(union Cell1787_4 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL1787_4 : KIND_EMPTY1787_4;
}
int gcd1787_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1787_5(int value) {
    return gcd1787_5(value, 68) * 88;
}
