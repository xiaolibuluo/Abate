
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2199_0(int value, ...) {
    return value + 74;
}
#else
int variant2199_0(int value) {
    return value + 74;
}
#endif
enum Kind2199_1 { KIND_EMPTY2199_1 = 0, KIND_FULL2199_1 = 81 };

union Cell2199_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2199_1(union Cell2199_1 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL2199_1 : KIND_EMPTY2199_1;
}
#define MASK2199_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2199_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2199_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
