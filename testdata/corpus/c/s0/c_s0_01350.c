
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind925_0 { KIND_EMPTY925_0 = 0, KIND_FULL925_0 = 99 };

union Cell925_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind925_0(union Cell925_0 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL925_0 : KIND_EMPTY925_0;
}
