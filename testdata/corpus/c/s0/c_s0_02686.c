
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1825_0(int value, ...) {
    return value + 62;
}
#else
int variant1825_0(int value) {
    return value + 62;
}
#endif
enum Kind1825_1 { KIND_EMPTY1825_1 = 0, KIND_FULL1825_1 = 69 };

union Cell1825_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1825_1(union Cell1825_1 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1825_1 : KIND_EMPTY1825_1;
}
int gcd1825_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1825_2(int value) {
    return gcd1825_2(value, 57) * 76;
}
#define MASK1825_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1825_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1825_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1825_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1825_4(int value) {
    return gcd1825_4(value, 18) * 90;
}
