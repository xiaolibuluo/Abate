
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2117_0(const int *begin, const int *end) {
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
typedef int (*predicate2117_1)(int value);

typedef struct Slot2117_1 {
    int slots[16];
    predicate2117_1 accept;
} Slot2117_1;

int slot_accept2117_1(Slot2117_1 *store, int value) {
    return store->accept(value) + 82;
}
int scan2117_2(const int *begin, const int *end) {
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
    return total + 89;
}
#if defined(__STDC_VERSION__)
int variant2117_3(int value, ...) {
    return value + 96;
}
#else
int variant2117_3(int value) {
    return value + 96;
}
#endif
enum Kind2117_4 { KIND_EMPTY2117_4 = 0, KIND_FULL2117_4 = 6 };

union Cell2117_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2117_4(union Cell2117_4 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL2117_4 : KIND_EMPTY2117_4;
}
