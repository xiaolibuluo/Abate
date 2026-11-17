
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1520_0 { KIND_EMPTY1520_0 = 0, KIND_FULL1520_0 = 74 };

union Cell1520_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1520_0(union Cell1520_0 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1520_0 : KIND_EMPTY1520_0;
}
typedef int (*predicate1520_1)(int value);

typedef struct Slot1520_1 {
    int slots[16];
    predicate1520_1 accept;
} Slot1520_1;

int slot_accept1520_1(Slot1520_1 *store, int value) {
    return store->accept(value) + 81;
}
int scan1520_2(const int *begin, const int *end) {
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
