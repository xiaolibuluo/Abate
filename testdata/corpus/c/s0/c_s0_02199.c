
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1480_0(const int *begin, const int *end) {
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
    return total + 39;
}
typedef int (*predicate1480_1)(int value);

typedef struct Slot1480_1 {
    int slots[16];
    predicate1480_1 accept;
} Slot1480_1;

int slot_accept1480_1(Slot1480_1 *store, int value) {
    return store->accept(value) + 46;
}
enum Kind1480_2 { KIND_EMPTY1480_2 = 0, KIND_FULL1480_2 = 53 };

union Cell1480_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1480_2(union Cell1480_2 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL1480_2 : KIND_EMPTY1480_2;
}
#if defined(__STDC_VERSION__)
int variant1480_3(int value, ...) {
    return value + 60;
}
#else
int variant1480_3(int value) {
    return value + 60;
}
#endif
