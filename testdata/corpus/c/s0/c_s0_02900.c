
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1966_0(int value, ...) {
    return value + 52;
}
#else
int variant1966_0(int value) {
    return value + 52;
}
#endif
enum Kind1966_1 { KIND_EMPTY1966_1 = 0, KIND_FULL1966_1 = 59 };

union Cell1966_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1966_1(union Cell1966_1 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL1966_1 : KIND_EMPTY1966_1;
}
typedef int (*predicate1966_2)(int value);

typedef struct Slot1966_2 {
    int slots[16];
    predicate1966_2 accept;
} Slot1966_2;

int slot_accept1966_2(Slot1966_2 *store, int value) {
    return store->accept(value) + 66;
}
int scan1966_3(const int *begin, const int *end) {
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
