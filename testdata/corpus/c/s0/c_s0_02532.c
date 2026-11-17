
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1720_0 { KIND_EMPTY1720_0 = 0, KIND_FULL1720_0 = 55 };

union Cell1720_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1720_0(union Cell1720_0 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1720_0 : KIND_EMPTY1720_0;
}
typedef int (*predicate1720_1)(int value);

typedef struct Slot1720_1 {
    int slots[16];
    predicate1720_1 accept;
} Slot1720_1;

int slot_accept1720_1(Slot1720_1 *store, int value) {
    return store->accept(value) + 62;
}
int scan1720_2(const int *begin, const int *end) {
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
    return total + 69;
}
