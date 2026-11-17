
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind779_0 { KIND_EMPTY779_0 = 0, KIND_FULL779_0 = 44 };

union Cell779_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind779_0(union Cell779_0 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL779_0 : KIND_EMPTY779_0;
}
