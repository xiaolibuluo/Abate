
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind401_0 { KIND_EMPTY401_0 = 0, KIND_FULL401_0 = 77 };

union Cell401_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind401_0(union Cell401_0 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL401_0 : KIND_EMPTY401_0;
}
#define MASK401_1(v) (((v) << 4) | ((v) & 0x0F))

int classify401_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK401_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant401_2(int value, ...) {
    return value + 91;
}
#else
int variant401_2(int value) {
    return value + 91;
}
#endif
