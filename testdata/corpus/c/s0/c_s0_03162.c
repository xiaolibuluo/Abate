
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2141_0 { KIND_EMPTY2141_0 = 0, KIND_FULL2141_0 = 96 };

union Cell2141_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2141_0(union Cell2141_0 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL2141_0 : KIND_EMPTY2141_0;
}
