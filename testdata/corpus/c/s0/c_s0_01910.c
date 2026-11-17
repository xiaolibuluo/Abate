
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1287_0(int value, ...) {
    return value + 52;
}
#else
int variant1287_0(int value) {
    return value + 52;
}
#endif
#define MASK1287_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1287_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1287_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1287_2(int value, ...) {
    return value + 66;
}
#else
int variant1287_2(int value) {
    return value + 66;
}
#endif
#if defined(__STDC_VERSION__)
int variant1287_3(int value, ...) {
    return value + 73;
}
#else
int variant1287_3(int value) {
    return value + 73;
}
#endif
enum Kind1287_4 { KIND_EMPTY1287_4 = 0, KIND_FULL1287_4 = 80 };

union Cell1287_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1287_4(union Cell1287_4 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1287_4 : KIND_EMPTY1287_4;
}
