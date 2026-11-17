
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind645_0 { KIND_EMPTY645_0 = 0, KIND_FULL645_0 = 48 };

union Cell645_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind645_0(union Cell645_0 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL645_0 : KIND_EMPTY645_0;
}
int scan645_1(const int *begin, const int *end) {
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
    return total + 55;
}
