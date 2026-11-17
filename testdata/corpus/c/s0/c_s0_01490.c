
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1012_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1012_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1012_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1012_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1012_1(int value) {
    return gcd1012_1(value, 28) * 73;
}
#if defined(__STDC_VERSION__)
int variant1012_2(int value, ...) {
    return value + 80;
}
#else
int variant1012_2(int value) {
    return value + 80;
}
#endif
#if defined(__STDC_VERSION__)
int variant1012_3(int value, ...) {
    return value + 87;
}
#else
int variant1012_3(int value) {
    return value + 87;
}
#endif
#define MASK1012_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1012_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1012_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1012_5 { KIND_EMPTY1012_5 = 0, KIND_FULL1012_5 = 101 };

union Cell1012_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1012_5(union Cell1012_5 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1012_5 : KIND_EMPTY1012_5;
}
enum Kind1012_6 { KIND_EMPTY1012_6 = 0, KIND_FULL1012_6 = 11 };

union Cell1012_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1012_6(union Cell1012_6 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1012_6 : KIND_EMPTY1012_6;
}
