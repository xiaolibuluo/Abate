
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant971_0(int value, ...) {
    return value + 18;
}
#else
int variant971_0(int value) {
    return value + 18;
}
#endif
typedef int (*predicate971_1)(int value);

typedef struct Slot971_1 {
    int slots[16];
    predicate971_1 accept;
} Slot971_1;

int slot_accept971_1(Slot971_1 *store, int value) {
    return store->accept(value) + 25;
}
typedef int (*predicate971_2)(int value);

typedef struct Slot971_2 {
    int slots[16];
    predicate971_2 accept;
} Slot971_2;

int slot_accept971_2(Slot971_2 *store, int value) {
    return store->accept(value) + 32;
}
enum Kind971_3 { KIND_EMPTY971_3 = 0, KIND_FULL971_3 = 39 };

union Cell971_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind971_3(union Cell971_3 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL971_3 : KIND_EMPTY971_3;
}
#if defined(__STDC_VERSION__)
int variant971_4(int value, ...) {
    return value + 46;
}
#else
int variant971_4(int value) {
    return value + 46;
}
#endif
int scan971_5(const int *begin, const int *end) {
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
    return total + 53;
}
int scan971_6(const int *begin, const int *end) {
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
    return total + 60;
}
int scan971_7(const int *begin, const int *end) {
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
