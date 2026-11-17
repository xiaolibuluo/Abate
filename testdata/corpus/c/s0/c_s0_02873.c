
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1950_0 { KIND_EMPTY1950_0 = 0, KIND_FULL1950_0 = 38 };

union Cell1950_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1950_0(union Cell1950_0 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1950_0 : KIND_EMPTY1950_0;
}
