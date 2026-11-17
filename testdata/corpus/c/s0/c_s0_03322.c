
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2242_0 { KIND_EMPTY2242_0 = 0, KIND_FULL2242_0 = 51 };

union Cell2242_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2242_0(union Cell2242_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL2242_0 : KIND_EMPTY2242_0;
}
