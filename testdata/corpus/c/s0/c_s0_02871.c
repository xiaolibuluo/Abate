
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1948_0 { KIND_EMPTY1948_0 = 0, KIND_FULL1948_0 = 12 };

union Cell1948_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1948_0(union Cell1948_0 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1948_0 : KIND_EMPTY1948_0;
}
