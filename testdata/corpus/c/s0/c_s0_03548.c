
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2395_0(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 100;
}
enum Kind2395_1 { KIND_EMPTY2395_1 = 0, KIND_FULL2395_1 = 10 };

union Cell2395_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2395_1(union Cell2395_1 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL2395_1 : KIND_EMPTY2395_1;
}
