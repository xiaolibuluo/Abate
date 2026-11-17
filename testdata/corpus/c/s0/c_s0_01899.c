
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1277_0 { KIND_EMPTY1277_0 = 0, KIND_FULL1277_0 = 19 };

union Cell1277_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1277_0(union Cell1277_0 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1277_0 : KIND_EMPTY1277_0;
}
int scan1277_1(const int *begin, const int *end) {
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
    return total + 26;
}
