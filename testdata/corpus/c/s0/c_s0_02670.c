
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1811_0 { KIND_EMPTY1811_0 = 0, KIND_FULL1811_0 = 74 };

union Cell1811_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1811_0(union Cell1811_0 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1811_0 : KIND_EMPTY1811_0;
}
#if defined(__STDC_VERSION__)
int variant1811_1(int value, ...) {
    return value + 81;
}
#else
int variant1811_1(int value) {
    return value + 81;
}
#endif
enum Kind1811_2 { KIND_EMPTY1811_2 = 0, KIND_FULL1811_2 = 88 };

union Cell1811_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1811_2(union Cell1811_2 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL1811_2 : KIND_EMPTY1811_2;
}
