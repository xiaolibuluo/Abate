
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind833_0 { KIND_EMPTY833_0 = 0, KIND_FULL833_0 = 67 };

union Cell833_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind833_0(union Cell833_0 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL833_0 : KIND_EMPTY833_0;
}
int scan833_1(const int *begin, const int *end) {
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
    return total + 74;
}
