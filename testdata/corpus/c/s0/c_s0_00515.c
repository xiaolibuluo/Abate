
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind357_0 { KIND_EMPTY357_0 = 0, KIND_FULL357_0 = 87 };

union Cell357_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind357_0(union Cell357_0 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL357_0 : KIND_EMPTY357_0;
}
int scan357_1(const int *begin, const int *end) {
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
    return total + 94;
}
#if defined(__STDC_VERSION__)
int variant357_2(int value, ...) {
    return value + 101;
}
#else
int variant357_2(int value) {
    return value + 101;
}
#endif
