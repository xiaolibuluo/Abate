
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1471_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1471_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1471_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1471_1(int value, ...) {
    return value + 26;
}
#else
int variant1471_1(int value) {
    return value + 26;
}
#endif
enum Kind1471_2 { KIND_EMPTY1471_2 = 0, KIND_FULL1471_2 = 33 };

union Cell1471_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1471_2(union Cell1471_2 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1471_2 : KIND_EMPTY1471_2;
}
