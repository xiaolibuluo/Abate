
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1952_0 { KIND_EMPTY1952_0 = 0, KIND_FULL1952_0 = 64 };

union Cell1952_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1952_0(union Cell1952_0 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL1952_0 : KIND_EMPTY1952_0;
}
enum Kind1952_1 { KIND_EMPTY1952_1 = 0, KIND_FULL1952_1 = 71 };

union Cell1952_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1952_1(union Cell1952_1 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL1952_1 : KIND_EMPTY1952_1;
}
