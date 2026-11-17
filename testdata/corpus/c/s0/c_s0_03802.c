
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2571_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2571_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2571_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2571_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2571_1(int value) {
    return gcd2571_1(value, 22) * 67;
}
#define MASK2571_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2571_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2571_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2571_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2571_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2571_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2571_4(const int *begin, const int *end) {
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
    return total + 88;
}
int scan2571_5(const int *begin, const int *end) {
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
    return total + 95;
}
enum Kind2571_6 { KIND_EMPTY2571_6 = 0, KIND_FULL2571_6 = 5 };

union Cell2571_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2571_6(union Cell2571_6 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL2571_6 : KIND_EMPTY2571_6;
}
