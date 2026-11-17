
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1246_0(const int *begin, const int *end) {
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
    return total + 101;
}
#define MASK1246_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1246_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1246_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1246_2(const int *begin, const int *end) {
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
    return total + 18;
}
enum Kind1246_3 { KIND_EMPTY1246_3 = 0, KIND_FULL1246_3 = 25 };

union Cell1246_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1246_3(union Cell1246_3 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL1246_3 : KIND_EMPTY1246_3;
}
int gcd1246_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1246_4(int value) {
    return gcd1246_4(value, 52) * 32;
}
