
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind692_0 { KIND_EMPTY692_0 = 0, KIND_FULL692_0 = 77 };

union Cell692_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind692_0(union Cell692_0 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL692_0 : KIND_EMPTY692_0;
}
enum Kind692_1 { KIND_EMPTY692_1 = 0, KIND_FULL692_1 = 84 };

union Cell692_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind692_1(union Cell692_1 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL692_1 : KIND_EMPTY692_1;
}
