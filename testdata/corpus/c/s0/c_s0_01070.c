
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind732_0 { KIND_EMPTY732_0 = 0, KIND_FULL732_0 = 15 };

union Cell732_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind732_0(union Cell732_0 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL732_0 : KIND_EMPTY732_0;
}
