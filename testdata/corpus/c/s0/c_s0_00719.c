
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant495_0(int value, ...) {
    return value + 38;
}
#else
int variant495_0(int value) {
    return value + 38;
}
#endif
enum Kind495_1 { KIND_EMPTY495_1 = 0, KIND_FULL495_1 = 45 };

union Cell495_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind495_1(union Cell495_1 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL495_1 : KIND_EMPTY495_1;
}
