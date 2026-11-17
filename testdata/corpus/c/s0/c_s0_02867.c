
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1947_0 { KIND_EMPTY1947_0 = 0, KIND_FULL1947_0 = 96 };

union Cell1947_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1947_0(union Cell1947_0 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1947_0 : KIND_EMPTY1947_0;
}
