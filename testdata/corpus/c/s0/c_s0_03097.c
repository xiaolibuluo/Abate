
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2095_0 { KIND_EMPTY2095_0 = 0, KIND_FULL2095_0 = 80 };

union Cell2095_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2095_0(union Cell2095_0 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL2095_0 : KIND_EMPTY2095_0;
}
