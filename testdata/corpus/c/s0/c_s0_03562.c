
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2404_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2404_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2404_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2404_1 { KIND_EMPTY2404_1 = 0, KIND_FULL2404_1 = 30 };

union Cell2404_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2404_1(union Cell2404_1 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL2404_1 : KIND_EMPTY2404_1;
}
#if defined(__STDC_VERSION__)
int variant2404_2(int value, ...) {
    return value + 37;
}
#else
int variant2404_2(int value) {
    return value + 37;
}
#endif
enum Kind2404_3 { KIND_EMPTY2404_3 = 0, KIND_FULL2404_3 = 44 };

union Cell2404_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2404_3(union Cell2404_3 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL2404_3 : KIND_EMPTY2404_3;
}
