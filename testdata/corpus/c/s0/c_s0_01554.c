
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1050_0 { KIND_EMPTY1050_0 = 0, KIND_FULL1050_0 = 75 };

union Cell1050_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1050_0(union Cell1050_0 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1050_0 : KIND_EMPTY1050_0;
}
#define MASK1050_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1050_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1050_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1050_2(int value, ...) {
    return value + 89;
}
#else
int variant1050_2(int value) {
    return value + 89;
}
#endif
int gcd1050_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1050_3(int value) {
    return gcd1050_3(value, 11) * 96;
}
int scan1050_4(const int *begin, const int *end) {
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
    return total + 6;
}
