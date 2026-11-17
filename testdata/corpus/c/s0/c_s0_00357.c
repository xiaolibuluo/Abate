
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind244_0 { KIND_EMPTY244_0 = 0, KIND_FULL244_0 = 73 };

union Cell244_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind244_0(union Cell244_0 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL244_0 : KIND_EMPTY244_0;
}
