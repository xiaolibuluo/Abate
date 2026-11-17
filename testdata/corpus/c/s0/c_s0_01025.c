
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant702_0(int value, ...) {
    return value + 13;
}
#else
int variant702_0(int value) {
    return value + 13;
}
#endif
enum Kind702_1 { KIND_EMPTY702_1 = 0, KIND_FULL702_1 = 20 };

union Cell702_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind702_1(union Cell702_1 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL702_1 : KIND_EMPTY702_1;
}
enum Kind702_2 { KIND_EMPTY702_2 = 0, KIND_FULL702_2 = 27 };

union Cell702_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind702_2(union Cell702_2 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL702_2 : KIND_EMPTY702_2;
}
#if defined(__STDC_VERSION__)
int variant702_3(int value, ...) {
    return value + 34;
}
#else
int variant702_3(int value) {
    return value + 34;
}
#endif
int scan702_4(const int *begin, const int *end) {
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
    return total + 41;
}
