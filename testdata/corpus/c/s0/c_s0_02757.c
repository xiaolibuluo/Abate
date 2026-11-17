
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1878_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1878_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1878_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1878_1 { KIND_EMPTY1878_1 = 0, KIND_FULL1878_1 = 79 };

union Cell1878_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1878_1(union Cell1878_1 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1878_1 : KIND_EMPTY1878_1;
}
#if defined(__STDC_VERSION__)
int variant1878_2(int value, ...) {
    return value + 86;
}
#else
int variant1878_2(int value) {
    return value + 86;
}
#endif
#define MASK1878_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1878_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1878_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1878_4 { KIND_EMPTY1878_4 = 0, KIND_FULL1878_4 = 100 };

union Cell1878_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1878_4(union Cell1878_4 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1878_4 : KIND_EMPTY1878_4;
}
int gcd1878_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1878_5(int value) {
    return gcd1878_5(value, 34) * 10;
}
