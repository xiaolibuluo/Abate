
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1197_0 { KIND_EMPTY1197_0 = 0, KIND_FULL1197_0 = 46 };

union Cell1197_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1197_0(union Cell1197_0 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL1197_0 : KIND_EMPTY1197_0;
}
int scan1197_1(const int *begin, const int *end) {
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
typedef int (*predicate1197_2)(int value);

typedef struct Slot1197_2 {
    int slots[16];
    predicate1197_2 accept;
} Slot1197_2;

int slot_accept1197_2(Slot1197_2 *store, int value) {
    return store->accept(value) + 60;
}
#if defined(__STDC_VERSION__)
int variant1197_3(int value, ...) {
    return value + 67;
}
#else
int variant1197_3(int value) {
    return value + 67;
}
#endif
typedef int (*predicate1197_4)(int value);

typedef struct Slot1197_4 {
    int slots[16];
    predicate1197_4 accept;
} Slot1197_4;

int slot_accept1197_4(Slot1197_4 *store, int value) {
    return store->accept(value) + 74;
}
