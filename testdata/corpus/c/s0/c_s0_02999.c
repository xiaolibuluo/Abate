
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2029_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2029_0(int value) {
    return gcd2029_0(value, 60) * 95;
}
#define MASK2029_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2029_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2029_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2029_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2029_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2029_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2029_3 { KIND_EMPTY2029_3 = 0, KIND_FULL2029_3 = 19 };

union Cell2029_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2029_3(union Cell2029_3 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL2029_3 : KIND_EMPTY2029_3;
}
enum Kind2029_4 { KIND_EMPTY2029_4 = 0, KIND_FULL2029_4 = 26 };

union Cell2029_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2029_4(union Cell2029_4 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL2029_4 : KIND_EMPTY2029_4;
}
#if defined(__STDC_VERSION__)
int variant2029_5(int value, ...) {
    return value + 33;
}
#else
int variant2029_5(int value) {
    return value + 33;
}
#endif
