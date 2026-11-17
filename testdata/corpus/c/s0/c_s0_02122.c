
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1429_0(const int *begin, const int *end) {
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
    return total + 55;
}
enum Kind1429_1 { KIND_EMPTY1429_1 = 0, KIND_FULL1429_1 = 62 };

union Cell1429_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1429_1(union Cell1429_1 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1429_1 : KIND_EMPTY1429_1;
}
#if defined(__STDC_VERSION__)
int variant1429_2(int value, ...) {
    return value + 69;
}
#else
int variant1429_2(int value) {
    return value + 69;
}
#endif
typedef int (*predicate1429_3)(int value);

typedef struct Slot1429_3 {
    int slots[16];
    predicate1429_3 accept;
} Slot1429_3;

int slot_accept1429_3(Slot1429_3 *store, int value) {
    return store->accept(value) + 76;
}
