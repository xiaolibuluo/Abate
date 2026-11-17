
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1419_0 { KIND_EMPTY1419_0 = 0, KIND_FULL1419_0 = 22 };

union Cell1419_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1419_0(union Cell1419_0 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1419_0 : KIND_EMPTY1419_0;
}
#if defined(__STDC_VERSION__)
int variant1419_1(int value, ...) {
    return value + 29;
}
#else
int variant1419_1(int value) {
    return value + 29;
}
#endif
