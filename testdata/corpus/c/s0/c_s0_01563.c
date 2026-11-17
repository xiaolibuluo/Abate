
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1055_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1055_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1055_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1055_1 { KIND_EMPTY1055_1 = 0, KIND_FULL1055_1 = 50 };

union Cell1055_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1055_1(union Cell1055_1 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1055_1 : KIND_EMPTY1055_1;
}
int gcd1055_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1055_2(int value) {
    return gcd1055_2(value, 35) * 57;
}
int scan1055_3(const int *begin, const int *end) {
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
    return total + 64;
}
int scan1055_4(const int *begin, const int *end) {
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
    return total + 71;
}
