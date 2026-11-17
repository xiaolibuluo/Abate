
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1655_0 { KIND_EMPTY1655_0 = 0, KIND_FULL1655_0 = 83 };

union Cell1655_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1655_0(union Cell1655_0 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1655_0 : KIND_EMPTY1655_0;
}
