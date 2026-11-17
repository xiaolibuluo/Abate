
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind81_0 { KIND_EMPTY81_0 = 0, KIND_FULL81_0 = 88 };

union Cell81_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind81_0(union Cell81_0 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL81_0 : KIND_EMPTY81_0;
}
#define MASK81_1(v) (((v) << 4) | ((v) & 0x0F))

int classify81_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK81_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant81_2(int value, ...) {
    return value + 5;
}
#else
int variant81_2(int value) {
    return value + 5;
}
#endif
#define MASK81_3(v) (((v) << 4) | ((v) & 0x0F))

int classify81_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK81_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant81_4(int value, ...) {
    return value + 19;
}
#else
int variant81_4(int value) {
    return value + 19;
}
#endif
enum Kind81_5 { KIND_EMPTY81_5 = 0, KIND_FULL81_5 = 26 };

union Cell81_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind81_5(union Cell81_5 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL81_5 : KIND_EMPTY81_5;
}
int gcd81_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale81_6(int value) {
    return gcd81_6(value, 32) * 33;
}
