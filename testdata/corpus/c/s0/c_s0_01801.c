
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1209_0(int value, ...) {
    return value + 8;
}
#else
int variant1209_0(int value) {
    return value + 8;
}
#endif
int gcd1209_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1209_1(int value) {
    return gcd1209_1(value, 65) * 15;
}
#define MASK1209_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1209_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1209_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1209_3(int value, ...) {
    return value + 29;
}
#else
int variant1209_3(int value) {
    return value + 29;
}
#endif
#define MASK1209_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1209_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1209_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1209_5(int value, ...) {
    return value + 43;
}
#else
int variant1209_5(int value) {
    return value + 43;
}
#endif
#define MASK1209_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1209_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1209_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1209_7 { KIND_EMPTY1209_7 = 0, KIND_FULL1209_7 = 57 };

union Cell1209_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1209_7(union Cell1209_7 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL1209_7 : KIND_EMPTY1209_7;
}
