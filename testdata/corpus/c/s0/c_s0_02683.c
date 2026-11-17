
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1822_0(int value, ...) {
    return value + 23;
}
#else
int variant1822_0(int value) {
    return value + 23;
}
#endif
int scan1822_1(const int *begin, const int *end) {
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
    return total + 30;
}
typedef int (*predicate1822_2)(int value);

typedef struct Slot1822_2 {
    int slots[16];
    predicate1822_2 accept;
} Slot1822_2;

int slot_accept1822_2(Slot1822_2 *store, int value) {
    return store->accept(value) + 37;
}
enum Kind1822_3 { KIND_EMPTY1822_3 = 0, KIND_FULL1822_3 = 44 };

union Cell1822_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1822_3(union Cell1822_3 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1822_3 : KIND_EMPTY1822_3;
}
