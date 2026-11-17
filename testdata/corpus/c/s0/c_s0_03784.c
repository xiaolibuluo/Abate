
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2560_0 { KIND_EMPTY2560_0 = 0, KIND_FULL2560_0 = 14 };

union Cell2560_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2560_0(union Cell2560_0 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL2560_0 : KIND_EMPTY2560_0;
}
#if defined(__STDC_VERSION__)
int variant2560_1(int value, ...) {
    return value + 21;
}
#else
int variant2560_1(int value) {
    return value + 21;
}
#endif
int gcd2560_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2560_2(int value) {
    return gcd2560_2(value, 17) * 28;
}
#define MASK2560_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2560_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2560_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2560_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2560_4(int value) {
    return gcd2560_4(value, 39) * 42;
}
int gcd2560_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2560_5(int value) {
    return gcd2560_5(value, 50) * 49;
}
