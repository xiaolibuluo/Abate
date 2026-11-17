
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd914_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale914_0(int value) {
    return gcd914_0(value, 63) * 53;
}
enum Kind914_1 { KIND_EMPTY914_1 = 0, KIND_FULL914_1 = 60 };

union Cell914_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind914_1(union Cell914_1 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL914_1 : KIND_EMPTY914_1;
}
#if defined(__STDC_VERSION__)
int variant914_2(int value, ...) {
    return value + 67;
}
#else
int variant914_2(int value) {
    return value + 67;
}
#endif
enum Kind914_3 { KIND_EMPTY914_3 = 0, KIND_FULL914_3 = 74 };

union Cell914_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind914_3(union Cell914_3 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL914_3 : KIND_EMPTY914_3;
}
#define MASK914_4(v) (((v) << 4) | ((v) & 0x0F))

int classify914_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK914_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd914_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale914_5(int value) {
    return gcd914_5(value, 57) * 88;
}
enum Kind914_6 { KIND_EMPTY914_6 = 0, KIND_FULL914_6 = 95 };

union Cell914_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind914_6(union Cell914_6 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL914_6 : KIND_EMPTY914_6;
}
