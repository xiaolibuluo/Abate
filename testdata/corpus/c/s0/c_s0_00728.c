
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind501_0 { KIND_EMPTY501_0 = 0, KIND_FULL501_0 = 19 };

union Cell501_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind501_0(union Cell501_0 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL501_0 : KIND_EMPTY501_0;
}
enum Kind501_1 { KIND_EMPTY501_1 = 0, KIND_FULL501_1 = 26 };

union Cell501_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind501_1(union Cell501_1 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL501_1 : KIND_EMPTY501_1;
}
enum Kind501_2 { KIND_EMPTY501_2 = 0, KIND_FULL501_2 = 33 };

union Cell501_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind501_2(union Cell501_2 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL501_2 : KIND_EMPTY501_2;
}
