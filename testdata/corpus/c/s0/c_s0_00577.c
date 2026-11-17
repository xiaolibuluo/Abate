
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind393_0 { KIND_EMPTY393_0 = 0, KIND_FULL393_0 = 70 };

union Cell393_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind393_0(union Cell393_0 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL393_0 : KIND_EMPTY393_0;
}
