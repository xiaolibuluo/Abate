
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1624_0(const int *begin, const int *end) {
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
    return total + 68;
}
enum Kind1624_1 { KIND_EMPTY1624_1 = 0, KIND_FULL1624_1 = 75 };

union Cell1624_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1624_1(union Cell1624_1 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1624_1 : KIND_EMPTY1624_1;
}
int scan1624_2(const int *begin, const int *end) {
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
    return total + 82;
}
