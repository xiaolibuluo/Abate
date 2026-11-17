
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind458_0 { KIND_EMPTY458_0 = 0, KIND_FULL458_0 = 42 };

union Cell458_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind458_0(union Cell458_0 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL458_0 : KIND_EMPTY458_0;
}
int scan458_1(const int *begin, const int *end) {
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
    return total + 49;
}
enum Kind458_2 { KIND_EMPTY458_2 = 0, KIND_FULL458_2 = 56 };

union Cell458_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind458_2(union Cell458_2 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL458_2 : KIND_EMPTY458_2;
}
enum Kind458_3 { KIND_EMPTY458_3 = 0, KIND_FULL458_3 = 63 };

union Cell458_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind458_3(union Cell458_3 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL458_3 : KIND_EMPTY458_3;
}
