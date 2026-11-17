
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2427_0 { KIND_EMPTY2427_0 = 0, KIND_FULL2427_0 = 31 };

union Cell2427_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2427_0(union Cell2427_0 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL2427_0 : KIND_EMPTY2427_0;
}
enum Kind2427_1 { KIND_EMPTY2427_1 = 0, KIND_FULL2427_1 = 38 };

union Cell2427_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2427_1(union Cell2427_1 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL2427_1 : KIND_EMPTY2427_1;
}
int gcd2427_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2427_2(int value) {
    return gcd2427_2(value, 62) * 45;
}
enum Kind2427_3 { KIND_EMPTY2427_3 = 0, KIND_FULL2427_3 = 52 };

union Cell2427_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2427_3(union Cell2427_3 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL2427_3 : KIND_EMPTY2427_3;
}
#if defined(__STDC_VERSION__)
int variant2427_4(int value, ...) {
    return value + 59;
}
#else
int variant2427_4(int value) {
    return value + 59;
}
#endif
int scan2427_5(const int *begin, const int *end) {
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
    return total + 66;
}
#define MASK2427_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2427_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2427_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
