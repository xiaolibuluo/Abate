
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1578_0 { KIND_EMPTY1578_0 = 0, KIND_FULL1578_0 = 52 };

union Cell1578_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1578_0(union Cell1578_0 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1578_0 : KIND_EMPTY1578_0;
}
int gcd1578_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1578_1(int value) {
    return gcd1578_1(value, 25) * 59;
}
#define MASK1578_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1578_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1578_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1578_3(int value, ...) {
    return value + 73;
}
#else
int variant1578_3(int value) {
    return value + 73;
}
#endif
