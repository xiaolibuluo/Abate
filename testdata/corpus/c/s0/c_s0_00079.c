
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind52_0 { KIND_EMPTY52_0 = 0, KIND_FULL52_0 = 99 };

union Cell52_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind52_0(union Cell52_0 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL52_0 : KIND_EMPTY52_0;
}
