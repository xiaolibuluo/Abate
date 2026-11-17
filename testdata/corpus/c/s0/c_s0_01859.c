
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1247_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1247_0(int value) {
    return gcd1247_0(value, 15) * 17;
}
int scan1247_1(const int *begin, const int *end) {
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
    return total + 24;
}
int scan1247_2(const int *begin, const int *end) {
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
    return total + 31;
}
#define MASK1247_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1247_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1247_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1247_4 { KIND_EMPTY1247_4 = 0, KIND_FULL1247_4 = 45 };

union Cell1247_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1247_4(union Cell1247_4 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1247_4 : KIND_EMPTY1247_4;
}
