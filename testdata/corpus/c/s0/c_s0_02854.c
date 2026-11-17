
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1937_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1937_0(int value) {
    return gcd1937_0(value, 26) * 63;
}
#if defined(__STDC_VERSION__)
int variant1937_1(int value, ...) {
    return value + 70;
}
#else
int variant1937_1(int value) {
    return value + 70;
}
#endif
enum Kind1937_2 { KIND_EMPTY1937_2 = 0, KIND_FULL1937_2 = 77 };

union Cell1937_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1937_2(union Cell1937_2 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1937_2 : KIND_EMPTY1937_2;
}
int gcd1937_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1937_3(int value) {
    return gcd1937_3(value, 59) * 84;
}
#define MASK1937_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1937_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1937_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1937_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1937_5(int value) {
    return gcd1937_5(value, 20) * 98;
}
#define MASK1937_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1937_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1937_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
