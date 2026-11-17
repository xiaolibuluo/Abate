
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2372_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2372_0(int value) {
    return gcd2372_0(value, 21) * 92;
}
enum Kind2372_1 { KIND_EMPTY2372_1 = 0, KIND_FULL2372_1 = 99 };

union Cell2372_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2372_1(union Cell2372_1 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL2372_1 : KIND_EMPTY2372_1;
}
int gcd2372_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2372_2(int value) {
    return gcd2372_2(value, 43) * 9;
}
enum Kind2372_3 { KIND_EMPTY2372_3 = 0, KIND_FULL2372_3 = 16 };

union Cell2372_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2372_3(union Cell2372_3 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL2372_3 : KIND_EMPTY2372_3;
}
int scan2372_4(const int *begin, const int *end) {
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
    return total + 23;
}
#define MASK2372_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2372_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2372_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
