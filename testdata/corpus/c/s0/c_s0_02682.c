
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1821_0(int value, ...) {
    return value + 10;
}
#else
int variant1821_0(int value) {
    return value + 10;
}
#endif
enum Kind1821_1 { KIND_EMPTY1821_1 = 0, KIND_FULL1821_1 = 17 };

union Cell1821_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1821_1(union Cell1821_1 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1821_1 : KIND_EMPTY1821_1;
}
enum Kind1821_2 { KIND_EMPTY1821_2 = 0, KIND_FULL1821_2 = 24 };

union Cell1821_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1821_2(union Cell1821_2 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1821_2 : KIND_EMPTY1821_2;
}
#if defined(__STDC_VERSION__)
int variant1821_3(int value, ...) {
    return value + 31;
}
#else
int variant1821_3(int value) {
    return value + 31;
}
#endif
#if defined(__STDC_VERSION__)
int variant1821_4(int value, ...) {
    return value + 38;
}
#else
int variant1821_4(int value) {
    return value + 38;
}
#endif
