
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1852_0 { KIND_EMPTY1852_0 = 0, KIND_FULL1852_0 = 25 };

union Cell1852_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1852_0(union Cell1852_0 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL1852_0 : KIND_EMPTY1852_0;
}
