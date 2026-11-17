
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind354_0 { KIND_EMPTY354_0 = 0, KIND_FULL354_0 = 48 };

union Cell354_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind354_0(union Cell354_0 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL354_0 : KIND_EMPTY354_0;
}
#define MASK354_1(v) (((v) << 4) | ((v) & 0x0F))

int classify354_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK354_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant354_2(int value, ...) {
    return value + 62;
}
#else
int variant354_2(int value) {
    return value + 62;
}
#endif
