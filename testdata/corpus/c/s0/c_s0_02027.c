
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1370_0 { KIND_EMPTY1370_0 = 0, KIND_FULL1370_0 = 64 };

union Cell1370_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1370_0(union Cell1370_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1370_0 : KIND_EMPTY1370_0;
}
enum Kind1370_1 { KIND_EMPTY1370_1 = 0, KIND_FULL1370_1 = 71 };

union Cell1370_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1370_1(union Cell1370_1 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL1370_1 : KIND_EMPTY1370_1;
}
