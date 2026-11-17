
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2016_0 { KIND_EMPTY2016_0 = 0, KIND_FULL2016_0 = 23 };

union Cell2016_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2016_0(union Cell2016_0 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL2016_0 : KIND_EMPTY2016_0;
}
int gcd2016_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2016_1(int value) {
    return gcd2016_1(value, 41) * 30;
}
#define MASK2016_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2016_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2016_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2016_3(const int *begin, const int *end) {
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
    return total + 44;
}
