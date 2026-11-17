
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind783_0 { KIND_EMPTY783_0 = 0, KIND_FULL783_0 = 96 };

union Cell783_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind783_0(union Cell783_0 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL783_0 : KIND_EMPTY783_0;
}
#if defined(__STDC_VERSION__)
int variant783_1(int value, ...) {
    return value + 6;
}
#else
int variant783_1(int value) {
    return value + 6;
}
#endif
