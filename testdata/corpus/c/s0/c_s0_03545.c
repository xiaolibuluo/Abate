
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2393_0(int value, ...) {
    return value + 74;
}
#else
int variant2393_0(int value) {
    return value + 74;
}
#endif
enum Kind2393_1 { KIND_EMPTY2393_1 = 0, KIND_FULL2393_1 = 81 };

union Cell2393_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2393_1(union Cell2393_1 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL2393_1 : KIND_EMPTY2393_1;
}
int scan2393_2(const int *begin, const int *end) {
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
    return total + 88;
}
