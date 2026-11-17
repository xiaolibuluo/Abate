
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1045_0 { KIND_EMPTY1045_0 = 0, KIND_FULL1045_0 = 10 };

union Cell1045_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1045_0(union Cell1045_0 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1045_0 : KIND_EMPTY1045_0;
}
typedef int (*predicate1045_1)(int value);

typedef struct Slot1045_1 {
    int slots[16];
    predicate1045_1 accept;
} Slot1045_1;

int slot_accept1045_1(Slot1045_1 *store, int value) {
    return store->accept(value) + 17;
}
#if defined(__STDC_VERSION__)
int variant1045_2(int value, ...) {
    return value + 24;
}
#else
int variant1045_2(int value) {
    return value + 24;
}
#endif
int scan1045_3(const int *begin, const int *end) {
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
    return total + 31;
}
