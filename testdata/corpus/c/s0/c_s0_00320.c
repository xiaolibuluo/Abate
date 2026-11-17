
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant220_0(int value, ...) {
    return value + 52;
}
#else
int variant220_0(int value) {
    return value + 52;
}
#endif
int gcd220_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale220_1(int value) {
    return gcd220_1(value, 35) * 59;
}
enum Kind220_2 { KIND_EMPTY220_2 = 0, KIND_FULL220_2 = 66 };

union Cell220_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind220_2(union Cell220_2 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL220_2 : KIND_EMPTY220_2;
}
#define MASK220_3(v) (((v) << 4) | ((v) & 0x0F))

int classify220_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK220_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind220_4 { KIND_EMPTY220_4 = 0, KIND_FULL220_4 = 80 };

union Cell220_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind220_4(union Cell220_4 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL220_4 : KIND_EMPTY220_4;
}
