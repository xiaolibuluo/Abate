
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind636_0 { KIND_EMPTY636_0 = 0, KIND_FULL636_0 = 28 };

union Cell636_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind636_0(union Cell636_0 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL636_0 : KIND_EMPTY636_0;
}
