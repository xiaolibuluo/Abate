
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan660_0(const int *begin, const int *end) {
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
    return total + 49;
}
int gcd660_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale660_1(int value) {
    return gcd660_1(value, 65) * 56;
}
enum Kind660_2 { KIND_EMPTY660_2 = 0, KIND_FULL660_2 = 63 };

union Cell660_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind660_2(union Cell660_2 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL660_2 : KIND_EMPTY660_2;
}
int gcd660_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale660_3(int value) {
    return gcd660_3(value, 26) * 70;
}
int scan660_4(const int *begin, const int *end) {
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
    return total + 77;
}
