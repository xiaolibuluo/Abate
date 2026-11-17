
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2539_0 { KIND_EMPTY2539_0 = 0, KIND_FULL2539_0 = 32 };

union Cell2539_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2539_0(union Cell2539_0 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL2539_0 : KIND_EMPTY2539_0;
}
int scan2539_1(const int *begin, const int *end) {
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
    return total + 39;
}
