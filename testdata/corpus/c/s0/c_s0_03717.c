
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2513_0(int value, ...) {
    return value + 82;
}
#else
int variant2513_0(int value) {
    return value + 82;
}
#endif
#if defined(__STDC_VERSION__)
int variant2513_1(int value, ...) {
    return value + 89;
}
#else
int variant2513_1(int value) {
    return value + 89;
}
#endif
int gcd2513_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2513_2(int value) {
    return gcd2513_2(value, 54) * 96;
}
#if defined(__STDC_VERSION__)
int variant2513_3(int value, ...) {
    return value + 6;
}
#else
int variant2513_3(int value) {
    return value + 6;
}
#endif
enum Kind2513_4 { KIND_EMPTY2513_4 = 0, KIND_FULL2513_4 = 13 };

union Cell2513_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2513_4(union Cell2513_4 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL2513_4 : KIND_EMPTY2513_4;
}
#if defined(__STDC_VERSION__)
int variant2513_5(int value, ...) {
    return value + 20;
}
#else
int variant2513_5(int value) {
    return value + 20;
}
#endif
#define MASK2513_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2513_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2513_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
