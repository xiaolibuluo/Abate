
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1574_0 { KIND_EMPTY1574_0 = 0, KIND_FULL1574_0 = 97 };

union Cell1574_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1574_0(union Cell1574_0 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1574_0 : KIND_EMPTY1574_0;
}
enum Kind1574_1 { KIND_EMPTY1574_1 = 0, KIND_FULL1574_1 = 7 };

union Cell1574_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1574_1(union Cell1574_1 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1574_1 : KIND_EMPTY1574_1;
}
int scan1574_2(const int *begin, const int *end) {
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
    return total + 14;
}
