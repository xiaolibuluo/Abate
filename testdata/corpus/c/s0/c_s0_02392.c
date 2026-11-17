
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1621_0)(int value);

typedef struct Slot1621_0 {
    int slots[16];
    predicate1621_0 accept;
} Slot1621_0;

int slot_accept1621_0(Slot1621_0 *store, int value) {
    return store->accept(value) + 29;
}
int scan1621_1(const int *begin, const int *end) {
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
    return total + 36;
}
enum Kind1621_2 { KIND_EMPTY1621_2 = 0, KIND_FULL1621_2 = 43 };

union Cell1621_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1621_2(union Cell1621_2 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1621_2 : KIND_EMPTY1621_2;
}
