
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind730_0 { KIND_EMPTY730_0 = 0, KIND_FULL730_0 = 86 };

union Cell730_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind730_0(union Cell730_0 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL730_0 : KIND_EMPTY730_0;
}
