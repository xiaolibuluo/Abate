
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1609_0 { KIND_EMPTY1609_0 = 0, KIND_FULL1609_0 = 67 };

union Cell1609_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1609_0(union Cell1609_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1609_0 : KIND_EMPTY1609_0;
}
