
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2083_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2083_0(int value) {
    return gcd2083_0(value, 11) * 21;
}
enum Kind2083_1 { KIND_EMPTY2083_1 = 0, KIND_FULL2083_1 = 28 };

union Cell2083_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2083_1(union Cell2083_1 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL2083_1 : KIND_EMPTY2083_1;
}
#define MASK2083_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2083_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2083_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2083_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2083_3(int value) {
    return gcd2083_3(value, 44) * 42;
}
int gcd2083_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2083_4(int value) {
    return gcd2083_4(value, 55) * 49;
}
int scan2083_5(const int *begin, const int *end) {
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
    return total + 56;
}
int scan2083_6(const int *begin, const int *end) {
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
    return total + 63;
}
