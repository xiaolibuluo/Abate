
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind637_0 { KIND_EMPTY637_0 = 0, KIND_FULL637_0 = 41 };

union Cell637_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind637_0(union Cell637_0 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL637_0 : KIND_EMPTY637_0;
}
#if defined(__STDC_VERSION__)
int variant637_1(int value, ...) {
    return value + 48;
}
#else
int variant637_1(int value) {
    return value + 48;
}
#endif
