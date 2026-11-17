
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1737_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1737_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1737_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1737_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1737_1(int value) {
    return gcd1737_1(value, 40) * 89;
}
enum Kind1737_2 { KIND_EMPTY1737_2 = 0, KIND_FULL1737_2 = 96 };

union Cell1737_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1737_2(union Cell1737_2 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1737_2 : KIND_EMPTY1737_2;
}
int scan1737_3(const int *begin, const int *end) {
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
int gcd1737_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1737_4(int value) {
    return gcd1737_4(value, 12) * 13;
}
#if defined(__STDC_VERSION__)
int variant1737_5(int value, ...) {
    return value + 20;
}
#else
int variant1737_5(int value) {
    return value + 20;
}
#endif
