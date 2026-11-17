
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2587_0)(int value);

typedef struct Slot2587_0 {
    int slots[16];
    predicate2587_0 accept;
} Slot2587_0;

int slot_accept2587_0(Slot2587_0 *store, int value) {
    return store->accept(value) + 74;
}
int scan2587_1(const int *begin, const int *end) {
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
    return total + 81;
}
