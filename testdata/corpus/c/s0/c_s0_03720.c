
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2516_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2516_0(int value) {
    return gcd2516_0(value, 53) * 24;
}
#define MASK2516_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2516_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2516_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2516_2(int value, ...) {
    return value + 38;
}
#else
int variant2516_2(int value) {
    return value + 38;
}
#endif
enum Kind2516_3 { KIND_EMPTY2516_3 = 0, KIND_FULL2516_3 = 45 };

union Cell2516_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2516_3(union Cell2516_3 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL2516_3 : KIND_EMPTY2516_3;
}
int scan2516_4(const int *begin, const int *end) {
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
    return total + 52;
}
enum Kind2516_5 { KIND_EMPTY2516_5 = 0, KIND_FULL2516_5 = 59 };

union Cell2516_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2516_5(union Cell2516_5 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL2516_5 : KIND_EMPTY2516_5;
}
