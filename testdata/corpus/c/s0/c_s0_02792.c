
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1897_0 { KIND_EMPTY1897_0 = 0, KIND_FULL1897_0 = 28 };

union Cell1897_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1897_0(union Cell1897_0 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1897_0 : KIND_EMPTY1897_0;
}
#if defined(__STDC_VERSION__)
int variant1897_1(int value, ...) {
    return value + 35;
}
#else
int variant1897_1(int value) {
    return value + 35;
}
#endif
#if defined(__STDC_VERSION__)
int variant1897_2(int value, ...) {
    return value + 42;
}
#else
int variant1897_2(int value) {
    return value + 42;
}
#endif
#if defined(__STDC_VERSION__)
int variant1897_3(int value, ...) {
    return value + 49;
}
#else
int variant1897_3(int value) {
    return value + 49;
}
#endif
int gcd1897_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1897_4(int value) {
    return gcd1897_4(value, 34) * 56;
}
#define MASK1897_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1897_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1897_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1897_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1897_6(int value) {
    return gcd1897_6(value, 56) * 70;
}
#define MASK1897_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1897_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1897_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1897_8(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1897_8(int value) {
    return gcd1897_8(value, 17) * 84;
}
