
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1295_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1295_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1295_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1295_1 { KIND_EMPTY1295_1 = 0, KIND_FULL1295_1 = 66 };

union Cell1295_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1295_1(union Cell1295_1 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL1295_1 : KIND_EMPTY1295_1;
}
#if defined(__STDC_VERSION__)
int variant1295_2(int value, ...) {
    return value + 73;
}
#else
int variant1295_2(int value) {
    return value + 73;
}
#endif
int gcd1295_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1295_3(int value) {
    return gcd1295_3(value, 18) * 80;
}
int scan1295_4(const int *begin, const int *end) {
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
    return total + 87;
}
int gcd1295_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1295_5(int value) {
    return gcd1295_5(value, 40) * 94;
}
