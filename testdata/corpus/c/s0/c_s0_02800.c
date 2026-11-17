
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1905_0 { KIND_EMPTY1905_0 = 0, KIND_FULL1905_0 = 35 };

union Cell1905_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1905_0(union Cell1905_0 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1905_0 : KIND_EMPTY1905_0;
}
enum Kind1905_1 { KIND_EMPTY1905_1 = 0, KIND_FULL1905_1 = 42 };

union Cell1905_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1905_1(union Cell1905_1 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL1905_1 : KIND_EMPTY1905_1;
}
