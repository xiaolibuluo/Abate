
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1612_0 { KIND_EMPTY1612_0 = 0, KIND_FULL1612_0 = 9 };

union Cell1612_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1612_0(union Cell1612_0 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1612_0 : KIND_EMPTY1612_0;
}
#if defined(__STDC_VERSION__)
int variant1612_1(int value, ...) {
    return value + 16;
}
#else
int variant1612_1(int value) {
    return value + 16;
}
#endif
