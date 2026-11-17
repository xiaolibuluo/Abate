
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind540_0 { KIND_EMPTY540_0 = 0, KIND_FULL540_0 = 41 };

union Cell540_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind540_0(union Cell540_0 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL540_0 : KIND_EMPTY540_0;
}
#if defined(__STDC_VERSION__)
int variant540_1(int value, ...) {
    return value + 48;
}
#else
int variant540_1(int value) {
    return value + 48;
}
#endif
