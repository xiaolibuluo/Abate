
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1853_0(int value, ...) {
    return value + 38;
}
#else
int variant1853_0(int value) {
    return value + 38;
}
#endif
enum Kind1853_1 { KIND_EMPTY1853_1 = 0, KIND_FULL1853_1 = 45 };

union Cell1853_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1853_1(union Cell1853_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1853_1 : KIND_EMPTY1853_1;
}
