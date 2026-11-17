
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1014_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1014_0(int value) {
    return gcd1014_0(value, 31) * 92;
}
#if defined(__STDC_VERSION__)
int variant1014_1(int value, ...) {
    return value + 99;
}
#else
int variant1014_1(int value) {
    return value + 99;
}
#endif
#define MASK1014_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1014_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1014_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1014_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1014_3(int value) {
    return gcd1014_3(value, 64) * 16;
}
enum Kind1014_4 { KIND_EMPTY1014_4 = 0, KIND_FULL1014_4 = 23 };

union Cell1014_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1014_4(union Cell1014_4 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1014_4 : KIND_EMPTY1014_4;
}
#if defined(__STDC_VERSION__)
int variant1014_5(int value, ...) {
    return value + 30;
}
#else
int variant1014_5(int value) {
    return value + 30;
}
#endif
int scan1014_6(const int *begin, const int *end) {
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
    return total + 37;
}
int gcd1014_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1014_7(int value) {
    return gcd1014_7(value, 47) * 44;
}
