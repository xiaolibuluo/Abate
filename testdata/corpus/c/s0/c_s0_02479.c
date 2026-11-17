
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1684_0(int value, ...) {
    return value + 72;
}
#else
int variant1684_0(int value) {
    return value + 72;
}
#endif
enum Kind1684_1 { KIND_EMPTY1684_1 = 0, KIND_FULL1684_1 = 79 };

union Cell1684_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1684_1(union Cell1684_1 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1684_1 : KIND_EMPTY1684_1;
}
enum Kind1684_2 { KIND_EMPTY1684_2 = 0, KIND_FULL1684_2 = 86 };

union Cell1684_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1684_2(union Cell1684_2 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1684_2 : KIND_EMPTY1684_2;
}
int gcd1684_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1684_3(int value) {
    return gcd1684_3(value, 57) * 93;
}
#define MASK1684_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1684_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1684_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1684_5(int value, ...) {
    return value + 10;
}
#else
int variant1684_5(int value) {
    return value + 10;
}
#endif
