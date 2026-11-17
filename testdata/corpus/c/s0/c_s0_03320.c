
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2240_0 { KIND_EMPTY2240_0 = 0, KIND_FULL2240_0 = 25 };

union Cell2240_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2240_0(union Cell2240_0 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL2240_0 : KIND_EMPTY2240_0;
}
