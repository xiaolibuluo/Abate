
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2423_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2423_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2423_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2423_1(const int *begin, const int *end) {
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
    return total + 83;
}
#if defined(__STDC_VERSION__)
int variant2423_2(int value, ...) {
    return value + 90;
}
#else
int variant2423_2(int value) {
    return value + 90;
}
#endif
#if defined(__STDC_VERSION__)
int variant2423_3(int value, ...) {
    return value + 97;
}
#else
int variant2423_3(int value) {
    return value + 97;
}
#endif
int gcd2423_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2423_4(int value) {
    return gcd2423_4(value, 56) * 7;
}
enum Kind2423_5 { KIND_EMPTY2423_5 = 0, KIND_FULL2423_5 = 14 };

union Cell2423_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2423_5(union Cell2423_5 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL2423_5 : KIND_EMPTY2423_5;
}
int scan2423_6(const int *begin, const int *end) {
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
    return total + 21;
}
