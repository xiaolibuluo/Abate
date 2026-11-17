
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2125_0(const int *begin, const int *end) {
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
    return total + 82;
}
enum Kind2125_1 { KIND_EMPTY2125_1 = 0, KIND_FULL2125_1 = 89 };

union Cell2125_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2125_1(union Cell2125_1 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL2125_1 : KIND_EMPTY2125_1;
}
#define MASK2125_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2125_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2125_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2125_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2125_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2125_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2125_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2125_4(int value) {
    return gcd2125_4(value, 44) * 13;
}
enum Kind2125_5 { KIND_EMPTY2125_5 = 0, KIND_FULL2125_5 = 20 };

union Cell2125_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2125_5(union Cell2125_5 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL2125_5 : KIND_EMPTY2125_5;
}
