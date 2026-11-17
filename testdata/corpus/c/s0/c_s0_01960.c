
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1323_0(int value, ...) {
    return value + 35;
}
#else
int variant1323_0(int value) {
    return value + 35;
}
#endif
enum Kind1323_1 { KIND_EMPTY1323_1 = 0, KIND_FULL1323_1 = 42 };

union Cell1323_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1323_1(union Cell1323_1 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL1323_1 : KIND_EMPTY1323_1;
}
enum Kind1323_2 { KIND_EMPTY1323_2 = 0, KIND_FULL1323_2 = 49 };

union Cell1323_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1323_2(union Cell1323_2 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL1323_2 : KIND_EMPTY1323_2;
}
