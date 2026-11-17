
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind500_0 { KIND_EMPTY500_0 = 0, KIND_FULL500_0 = 6 };

union Cell500_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind500_0(union Cell500_0 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL500_0 : KIND_EMPTY500_0;
}
enum Kind500_1 { KIND_EMPTY500_1 = 0, KIND_FULL500_1 = 13 };

union Cell500_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind500_1(union Cell500_1 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL500_1 : KIND_EMPTY500_1;
}
int scan500_2(const int *begin, const int *end) {
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
    return total + 20;
}
