
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1081_0 { KIND_EMPTY1081_0 = 0, KIND_FULL1081_0 = 90 };

union Cell1081_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1081_0(union Cell1081_0 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1081_0 : KIND_EMPTY1081_0;
}
int scan1081_1(const int *begin, const int *end) {
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
    return total + 97;
}
