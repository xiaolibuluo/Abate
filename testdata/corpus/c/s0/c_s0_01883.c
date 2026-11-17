
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1265_0 { KIND_EMPTY1265_0 = 0, KIND_FULL1265_0 = 57 };

union Cell1265_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1265_0(union Cell1265_0 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1265_0 : KIND_EMPTY1265_0;
}
