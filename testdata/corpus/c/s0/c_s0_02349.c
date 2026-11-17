
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1585_0(int value, ...) {
    return value + 46;
}
#else
int variant1585_0(int value) {
    return value + 46;
}
#endif
enum Kind1585_1 { KIND_EMPTY1585_1 = 0, KIND_FULL1585_1 = 53 };

union Cell1585_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1585_1(union Cell1585_1 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL1585_1 : KIND_EMPTY1585_1;
}
enum Kind1585_2 { KIND_EMPTY1585_2 = 0, KIND_FULL1585_2 = 60 };

union Cell1585_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1585_2(union Cell1585_2 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1585_2 : KIND_EMPTY1585_2;
}
#if defined(__STDC_VERSION__)
int variant1585_3(int value, ...) {
    return value + 67;
}
#else
int variant1585_3(int value) {
    return value + 67;
}
#endif
int gcd1585_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1585_4(int value) {
    return gcd1585_4(value, 46) * 74;
}
#define MASK1585_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1585_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1585_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
