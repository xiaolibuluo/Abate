
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1278_0(int value, ...) {
    return value + 32;
}
#else
int variant1278_0(int value) {
    return value + 32;
}
#endif
enum Kind1278_1 { KIND_EMPTY1278_1 = 0, KIND_FULL1278_1 = 39 };

union Cell1278_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1278_1(union Cell1278_1 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1278_1 : KIND_EMPTY1278_1;
}
enum Kind1278_2 { KIND_EMPTY1278_2 = 0, KIND_FULL1278_2 = 46 };

union Cell1278_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1278_2(union Cell1278_2 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL1278_2 : KIND_EMPTY1278_2;
}
