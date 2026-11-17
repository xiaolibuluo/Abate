
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1704_0 { KIND_EMPTY1704_0 = 0, KIND_FULL1704_0 = 41 };

union Cell1704_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1704_0(union Cell1704_0 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1704_0 : KIND_EMPTY1704_0;
}
