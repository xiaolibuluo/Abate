
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2534_0 { KIND_EMPTY2534_0 = 0, KIND_FULL2534_0 = 64 };

union Cell2534_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2534_0(union Cell2534_0 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL2534_0 : KIND_EMPTY2534_0;
}
