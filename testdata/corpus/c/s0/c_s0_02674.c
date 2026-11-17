
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1814_0 { KIND_EMPTY1814_0 = 0, KIND_FULL1814_0 = 16 };

union Cell1814_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1814_0(union Cell1814_0 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1814_0 : KIND_EMPTY1814_0;
}
int scan1814_1(const int *begin, const int *end) {
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
    return total + 23;
}
int scan1814_2(const int *begin, const int *end) {
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
    return total + 30;
}
