
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2467_0(int value, ...) {
    return value + 66;
}
#else
int variant2467_0(int value) {
    return value + 66;
}
#endif
int scan2467_1(const int *begin, const int *end) {
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
    return total + 73;
}
int scan2467_2(const int *begin, const int *end) {
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
    return total + 80;
}
enum Kind2467_3 { KIND_EMPTY2467_3 = 0, KIND_FULL2467_3 = 87 };

union Cell2467_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2467_3(union Cell2467_3 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL2467_3 : KIND_EMPTY2467_3;
}
typedef int (*predicate2467_4)(int value);

typedef struct Slot2467_4 {
    int slots[16];
    predicate2467_4 accept;
} Slot2467_4;

int slot_accept2467_4(Slot2467_4 *store, int value) {
    return store->accept(value) + 94;
}
#if defined(__STDC_VERSION__)
int variant2467_5(int value, ...) {
    return value + 101;
}
#else
int variant2467_5(int value) {
    return value + 101;
}
#endif
