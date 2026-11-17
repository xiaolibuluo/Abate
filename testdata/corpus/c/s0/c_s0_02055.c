
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1389_0 { KIND_EMPTY1389_0 = 0, KIND_FULL1389_0 = 20 };

union Cell1389_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1389_0(union Cell1389_0 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL1389_0 : KIND_EMPTY1389_0;
}
#define MASK1389_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1389_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1389_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1389_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1389_2(int value) {
    return gcd1389_2(value, 55) * 34;
}
enum Kind1389_3 { KIND_EMPTY1389_3 = 0, KIND_FULL1389_3 = 41 };

union Cell1389_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1389_3(union Cell1389_3 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1389_3 : KIND_EMPTY1389_3;
}
int scan1389_4(const int *begin, const int *end) {
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
    return total + 48;
}
