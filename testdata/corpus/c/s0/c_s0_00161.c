
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind110_0 { KIND_EMPTY110_0 = 0, KIND_FULL110_0 = 77 };

union Cell110_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind110_0(union Cell110_0 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL110_0 : KIND_EMPTY110_0;
}
int scan110_1(const int *begin, const int *end) {
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
    return total + 84;
}
