
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant829_0(int value, ...) {
    return value + 15;
}
#else
int variant829_0(int value) {
    return value + 15;
}
#endif
enum Kind829_1 { KIND_EMPTY829_1 = 0, KIND_FULL829_1 = 22 };

union Cell829_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind829_1(union Cell829_1 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL829_1 : KIND_EMPTY829_1;
}
