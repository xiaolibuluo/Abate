
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1316_0(int value, ...) {
    return value + 41;
}
#else
int variant1316_0(int value) {
    return value + 41;
}
#endif
enum Kind1316_1 { KIND_EMPTY1316_1 = 0, KIND_FULL1316_1 = 48 };

union Cell1316_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1316_1(union Cell1316_1 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL1316_1 : KIND_EMPTY1316_1;
}
