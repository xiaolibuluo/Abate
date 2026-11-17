
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2225_0(int value, ...) {
    return value + 24;
}
#else
int variant2225_0(int value) {
    return value + 24;
}
#endif
int gcd2225_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2225_1(int value) {
    return gcd2225_1(value, 40) * 31;
}
int gcd2225_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2225_2(int value) {
    return gcd2225_2(value, 51) * 38;
}
int scan2225_3(const int *begin, const int *end) {
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
enum Kind2225_4 { KIND_EMPTY2225_4 = 0, KIND_FULL2225_4 = 52 };

union Cell2225_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2225_4(union Cell2225_4 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL2225_4 : KIND_EMPTY2225_4;
}
#define MASK2225_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2225_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2225_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
