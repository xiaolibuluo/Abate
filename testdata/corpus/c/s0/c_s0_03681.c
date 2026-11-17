
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2488_0(const int *begin, const int *end) {
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
    return total + 48;
}
enum Kind2488_1 { KIND_EMPTY2488_1 = 0, KIND_FULL2488_1 = 55 };

union Cell2488_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2488_1(union Cell2488_1 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL2488_1 : KIND_EMPTY2488_1;
}
