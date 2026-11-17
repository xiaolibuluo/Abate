
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1349_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1349_0(int value) {
    return gcd1349_0(value, 58) * 82;
}
enum Kind1349_1 { KIND_EMPTY1349_1 = 0, KIND_FULL1349_1 = 89 };

union Cell1349_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1349_1(union Cell1349_1 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1349_1 : KIND_EMPTY1349_1;
}
enum Kind1349_2 { KIND_EMPTY1349_2 = 0, KIND_FULL1349_2 = 96 };

union Cell1349_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1349_2(union Cell1349_2 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1349_2 : KIND_EMPTY1349_2;
}
#define MASK1349_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1349_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1349_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1349_4(const int *begin, const int *end) {
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
    return total + 13;
}
int scan1349_5(const int *begin, const int *end) {
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
    return total + 20;
}
