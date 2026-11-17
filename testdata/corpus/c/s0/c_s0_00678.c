
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind464_0 { KIND_EMPTY464_0 = 0, KIND_FULL464_0 = 23 };

union Cell464_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind464_0(union Cell464_0 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL464_0 : KIND_EMPTY464_0;
}
int gcd464_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale464_1(int value) {
    return gcd464_1(value, 35) * 30;
}
enum Kind464_2 { KIND_EMPTY464_2 = 0, KIND_FULL464_2 = 37 };

union Cell464_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind464_2(union Cell464_2 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL464_2 : KIND_EMPTY464_2;
}
int scan464_3(const int *begin, const int *end) {
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
    return total + 44;
}
