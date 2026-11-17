
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind443_0 { KIND_EMPTY443_0 = 0, KIND_FULL443_0 = 41 };

union Cell443_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind443_0(union Cell443_0 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL443_0 : KIND_EMPTY443_0;
}
#if defined(__STDC_VERSION__)
int variant443_1(int value, ...) {
    return value + 48;
}
#else
int variant443_1(int value) {
    return value + 48;
}
#endif
