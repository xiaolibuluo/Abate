
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1575_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1575_0(int value) {
    return gcd1575_0(value, 54) * 13;
}
#if defined(__STDC_VERSION__)
int variant1575_1(int value, ...) {
    return value + 20;
}
#else
int variant1575_1(int value) {
    return value + 20;
}
#endif
#define MASK1575_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1575_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1575_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1575_3 { KIND_EMPTY1575_3 = 0, KIND_FULL1575_3 = 34 };

union Cell1575_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1575_3(union Cell1575_3 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL1575_3 : KIND_EMPTY1575_3;
}
