
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1452_0)(int value);

typedef struct Slot1452_0 {
    int slots[16];
    predicate1452_0 accept;
} Slot1452_0;

int slot_accept1452_0(Slot1452_0 *store, int value) {
    return store->accept(value) + 63;
}
int scan1452_1(const int *begin, const int *end) {
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
    return total + 70;
}
#if defined(__STDC_VERSION__)
int variant1452_2(int value, ...) {
    return value + 77;
}
#else
int variant1452_2(int value) {
    return value + 77;
}
#endif
#if defined(__STDC_VERSION__)
int variant1452_3(int value, ...) {
    return value + 84;
}
#else
int variant1452_3(int value) {
    return value + 84;
}
#endif
int scan1452_4(const int *begin, const int *end) {
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
    return total + 91;
}
#if defined(__STDC_VERSION__)
int variant1452_5(int value, ...) {
    return value + 98;
}
#else
int variant1452_5(int value) {
    return value + 98;
}
#endif
typedef int (*predicate1452_6)(int value);

typedef struct Slot1452_6 {
    int slots[16];
    predicate1452_6 accept;
} Slot1452_6;

int slot_accept1452_6(Slot1452_6 *store, int value) {
    return store->accept(value) + 8;
}
enum Kind1452_7 { KIND_EMPTY1452_7 = 0, KIND_FULL1452_7 = 15 };

union Cell1452_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1452_7(union Cell1452_7 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL1452_7 : KIND_EMPTY1452_7;
}
