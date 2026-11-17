
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2378_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2378_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2378_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2378_1(const int *begin, const int *end) {
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
enum Kind2378_2 { KIND_EMPTY2378_2 = 0, KIND_FULL2378_2 = 87 };

union Cell2378_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2378_2(union Cell2378_2 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL2378_2 : KIND_EMPTY2378_2;
}
int scan2378_3(const int *begin, const int *end) {
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
    return total + 94;
}
int gcd2378_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2378_4(int value) {
    return gcd2378_4(value, 46) * 101;
}
#define MASK2378_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2378_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2378_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2378_6(int value, ...) {
    return value + 18;
}
#else
int variant2378_6(int value) {
    return value + 18;
}
#endif
