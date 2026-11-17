
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1647_0(const int *begin, const int *end) {
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
    return total + 76;
}
typedef int (*predicate1647_1)(int value);

typedef struct Slot1647_1 {
    int slots[16];
    predicate1647_1 accept;
} Slot1647_1;

int slot_accept1647_1(Slot1647_1 *store, int value) {
    return store->accept(value) + 83;
}
#if defined(__STDC_VERSION__)
int variant1647_2(int value, ...) {
    return value + 90;
}
#else
int variant1647_2(int value) {
    return value + 90;
}
#endif
enum Kind1647_3 { KIND_EMPTY1647_3 = 0, KIND_FULL1647_3 = 97 };

union Cell1647_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1647_3(union Cell1647_3 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1647_3 : KIND_EMPTY1647_3;
}
int scan1647_4(const int *begin, const int *end) {
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
    return total + 7;
}
int scan1647_5(const int *begin, const int *end) {
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
    return total + 14;
}
int scan1647_6(const int *begin, const int *end) {
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
    return total + 21;
}
