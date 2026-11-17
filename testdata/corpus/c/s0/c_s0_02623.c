
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1777_0 { KIND_EMPTY1777_0 = 0, KIND_FULL1777_0 = 20 };

union Cell1777_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1777_0(union Cell1777_0 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1777_0 : KIND_EMPTY1777_0;
}
#define MASK1777_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1777_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1777_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1777_2(int value, ...) {
    return value + 34;
}
#else
int variant1777_2(int value) {
    return value + 34;
}
#endif
int gcd1777_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1777_3(int value) {
    return gcd1777_3(value, 37) * 41;
}
#define MASK1777_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1777_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1777_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
