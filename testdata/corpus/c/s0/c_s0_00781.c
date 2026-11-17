
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind539_0 { KIND_EMPTY539_0 = 0, KIND_FULL539_0 = 28 };

union Cell539_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind539_0(union Cell539_0 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL539_0 : KIND_EMPTY539_0;
}
