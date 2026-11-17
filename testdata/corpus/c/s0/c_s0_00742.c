
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind510_0 { KIND_EMPTY510_0 = 0, KIND_FULL510_0 = 39 };

union Cell510_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind510_0(union Cell510_0 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL510_0 : KIND_EMPTY510_0;
}
enum Kind510_1 { KIND_EMPTY510_1 = 0, KIND_FULL510_1 = 46 };

union Cell510_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind510_1(union Cell510_1 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL510_1 : KIND_EMPTY510_1;
}
enum Kind510_2 { KIND_EMPTY510_2 = 0, KIND_FULL510_2 = 53 };

union Cell510_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind510_2(union Cell510_2 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL510_2 : KIND_EMPTY510_2;
}
#if defined(__STDC_VERSION__)
int variant510_3(int value, ...) {
    return value + 60;
}
#else
int variant510_3(int value) {
    return value + 60;
}
#endif
int scan510_4(const int *begin, const int *end) {
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
    return total + 67;
}
