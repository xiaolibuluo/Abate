
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind309_0 { KIND_EMPTY309_0 = 0, KIND_FULL309_0 = 45 };

union Cell309_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind309_0(union Cell309_0 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL309_0 : KIND_EMPTY309_0;
}
enum Kind309_1 { KIND_EMPTY309_1 = 0, KIND_FULL309_1 = 52 };

union Cell309_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind309_1(union Cell309_1 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL309_1 : KIND_EMPTY309_1;
}
int scan309_2(const int *begin, const int *end) {
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
    return total + 59;
}
