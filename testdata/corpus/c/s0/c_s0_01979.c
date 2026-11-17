
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1336_0)(int value);

typedef struct Slot1336_0 {
    int slots[16];
    predicate1336_0 accept;
} Slot1336_0;

int slot_accept1336_0(Slot1336_0 *store, int value) {
    return store->accept(value) + 10;
}
typedef int (*predicate1336_1)(int value);

typedef struct Slot1336_1 {
    int slots[16];
    predicate1336_1 accept;
} Slot1336_1;

int slot_accept1336_1(Slot1336_1 *store, int value) {
    return store->accept(value) + 17;
}
int scan1336_2(const int *begin, const int *end) {
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
    return total + 24;
}
enum Kind1336_3 { KIND_EMPTY1336_3 = 0, KIND_FULL1336_3 = 31 };

union Cell1336_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1336_3(union Cell1336_3 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL1336_3 : KIND_EMPTY1336_3;
}
