
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1148_0(const int *begin, const int *end) {
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
typedef int (*predicate1148_1)(int value);

typedef struct Slot1148_1 {
    int slots[16];
    predicate1148_1 accept;
} Slot1148_1;

int slot_accept1148_1(Slot1148_1 *store, int value) {
    return store->accept(value) + 95;
}
#if defined(__STDC_VERSION__)
int variant1148_2(int value, ...) {
    return value + 5;
}
#else
int variant1148_2(int value) {
    return value + 5;
}
#endif
typedef int (*predicate1148_3)(int value);

typedef struct Slot1148_3 {
    int slots[16];
    predicate1148_3 accept;
} Slot1148_3;

int slot_accept1148_3(Slot1148_3 *store, int value) {
    return store->accept(value) + 12;
}
typedef int (*predicate1148_4)(int value);

typedef struct Slot1148_4 {
    int slots[16];
    predicate1148_4 accept;
} Slot1148_4;

int slot_accept1148_4(Slot1148_4 *store, int value) {
    return store->accept(value) + 19;
}
enum Kind1148_5 { KIND_EMPTY1148_5 = 0, KIND_FULL1148_5 = 26 };

union Cell1148_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1148_5(union Cell1148_5 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1148_5 : KIND_EMPTY1148_5;
}
