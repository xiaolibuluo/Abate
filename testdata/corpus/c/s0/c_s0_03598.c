
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2430_0 { KIND_EMPTY2430_0 = 0, KIND_FULL2430_0 = 70 };

union Cell2430_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2430_0(union Cell2430_0 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL2430_0 : KIND_EMPTY2430_0;
}
int scan2430_1(const int *begin, const int *end) {
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
int scan2430_2(const int *begin, const int *end) {
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
    return total + 84;
}
#define MASK2430_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2430_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2430_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2430_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2430_4(int value) {
    return gcd2430_4(value, 44) * 98;
}
enum Kind2430_5 { KIND_EMPTY2430_5 = 0, KIND_FULL2430_5 = 8 };

union Cell2430_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2430_5(union Cell2430_5 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL2430_5 : KIND_EMPTY2430_5;
}
#if defined(__STDC_VERSION__)
int variant2430_6(int value, ...) {
    return value + 15;
}
#else
int variant2430_6(int value) {
    return value + 15;
}
#endif
int scan2430_7(const int *begin, const int *end) {
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
    return total + 22;
}
