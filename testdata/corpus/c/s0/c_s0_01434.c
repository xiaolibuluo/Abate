
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind978_0 { KIND_EMPTY978_0 = 0, KIND_FULL978_0 = 12 };

union Cell978_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind978_0(union Cell978_0 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL978_0 : KIND_EMPTY978_0;
}
#if defined(__STDC_VERSION__)
int variant978_1(int value, ...) {
    return value + 19;
}
#else
int variant978_1(int value) {
    return value + 19;
}
#endif
