
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan396_0(const int *begin, const int *end) {
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
    return total + 12;
}
enum Kind396_1 { KIND_EMPTY396_1 = 0, KIND_FULL396_1 = 19 };

union Cell396_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind396_1(union Cell396_1 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL396_1 : KIND_EMPTY396_1;
}
