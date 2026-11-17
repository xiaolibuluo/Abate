
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1088_0(const int *begin, const int *end) {
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
enum Kind1088_1 { KIND_EMPTY1088_1 = 0, KIND_FULL1088_1 = 91 };

union Cell1088_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1088_1(union Cell1088_1 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1088_1 : KIND_EMPTY1088_1;
}
enum Kind1088_2 { KIND_EMPTY1088_2 = 0, KIND_FULL1088_2 = 98 };

union Cell1088_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1088_2(union Cell1088_2 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1088_2 : KIND_EMPTY1088_2;
}
