
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2505_0(const int *begin, const int *end) {
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
    return total + 75;
}
enum Kind2505_1 { KIND_EMPTY2505_1 = 0, KIND_FULL2505_1 = 82 };

union Cell2505_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2505_1(union Cell2505_1 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL2505_1 : KIND_EMPTY2505_1;
}
enum Kind2505_2 { KIND_EMPTY2505_2 = 0, KIND_FULL2505_2 = 89 };

union Cell2505_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2505_2(union Cell2505_2 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL2505_2 : KIND_EMPTY2505_2;
}
int scan2505_3(const int *begin, const int *end) {
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
    return total + 96;
}
