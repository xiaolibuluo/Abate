
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1184_0 { KIND_EMPTY1184_0 = 0, KIND_FULL1184_0 = 71 };

union Cell1184_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1184_0(union Cell1184_0 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL1184_0 : KIND_EMPTY1184_0;
}
typedef int (*predicate1184_1)(int value);

typedef struct Slot1184_1 {
    int slots[16];
    predicate1184_1 accept;
} Slot1184_1;

int slot_accept1184_1(Slot1184_1 *store, int value) {
    return store->accept(value) + 78;
}
int scan1184_2(const int *begin, const int *end) {
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
    return total + 85;
}
