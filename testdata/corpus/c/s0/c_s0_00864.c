
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind594_0 { KIND_EMPTY594_0 = 0, KIND_FULL594_0 = 64 };

union Cell594_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind594_0(union Cell594_0 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL594_0 : KIND_EMPTY594_0;
}
#if defined(__STDC_VERSION__)
int variant594_1(int value, ...) {
    return value + 71;
}
#else
int variant594_1(int value) {
    return value + 71;
}
#endif
#if defined(__STDC_VERSION__)
int variant594_2(int value, ...) {
    return value + 78;
}
#else
int variant594_2(int value) {
    return value + 78;
}
#endif
