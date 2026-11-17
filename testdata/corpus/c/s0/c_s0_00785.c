
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind542_0 { KIND_EMPTY542_0 = 0, KIND_FULL542_0 = 67 };

union Cell542_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind542_0(union Cell542_0 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL542_0 : KIND_EMPTY542_0;
}
int scan542_1(const int *begin, const int *end) {
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
