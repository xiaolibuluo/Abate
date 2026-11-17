
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1320_0 { KIND_EMPTY1320_0 = 0, KIND_FULL1320_0 = 93 };

union Cell1320_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1320_0(union Cell1320_0 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1320_0 : KIND_EMPTY1320_0;
}
#if defined(__STDC_VERSION__)
int variant1320_1(int value, ...) {
    return value + 100;
}
#else
int variant1320_1(int value) {
    return value + 100;
}
#endif
