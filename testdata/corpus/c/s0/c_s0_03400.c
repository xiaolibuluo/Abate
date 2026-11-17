
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2293_0 { KIND_EMPTY2293_0 = 0, KIND_FULL2293_0 = 35 };

union Cell2293_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2293_0(union Cell2293_0 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL2293_0 : KIND_EMPTY2293_0;
}
enum Kind2293_1 { KIND_EMPTY2293_1 = 0, KIND_FULL2293_1 = 42 };

union Cell2293_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2293_1(union Cell2293_1 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL2293_1 : KIND_EMPTY2293_1;
}
