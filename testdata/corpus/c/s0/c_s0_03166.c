
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2144_0 { KIND_EMPTY2144_0 = 0, KIND_FULL2144_0 = 38 };

union Cell2144_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2144_0(union Cell2144_0 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL2144_0 : KIND_EMPTY2144_0;
}
