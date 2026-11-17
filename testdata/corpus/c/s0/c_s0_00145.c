
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind99_0 { KIND_EMPTY99_0 = 0, KIND_FULL99_0 = 31 };

union Cell99_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind99_0(union Cell99_0 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL99_0 : KIND_EMPTY99_0;
}
