
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind108_0 { KIND_EMPTY108_0 = 0, KIND_FULL108_0 = 51 };

union Cell108_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind108_0(union Cell108_0 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL108_0 : KIND_EMPTY108_0;
}
enum Kind108_1 { KIND_EMPTY108_1 = 0, KIND_FULL108_1 = 58 };

union Cell108_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind108_1(union Cell108_1 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL108_1 : KIND_EMPTY108_1;
}
