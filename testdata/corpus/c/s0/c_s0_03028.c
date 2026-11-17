
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2047_0 { KIND_EMPTY2047_0 = 0, KIND_FULL2047_0 = 38 };

union Cell2047_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2047_0(union Cell2047_0 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL2047_0 : KIND_EMPTY2047_0;
}
