
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2416_0(const int *begin, const int *end) {
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
int gcd2416_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2416_1(int value) {
    return gcd2416_1(value, 35) * 89;
}
enum Kind2416_2 { KIND_EMPTY2416_2 = 0, KIND_FULL2416_2 = 96 };

union Cell2416_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2416_2(union Cell2416_2 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL2416_2 : KIND_EMPTY2416_2;
}
enum Kind2416_3 { KIND_EMPTY2416_3 = 0, KIND_FULL2416_3 = 6 };

union Cell2416_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2416_3(union Cell2416_3 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL2416_3 : KIND_EMPTY2416_3;
}
#define MASK2416_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2416_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2416_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
