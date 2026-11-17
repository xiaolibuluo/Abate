
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2351_0)(int value);

typedef struct Slot2351_0 {
    int slots[16];
    predicate2351_0 accept;
} Slot2351_0;

int slot_accept2351_0(Slot2351_0 *store, int value) {
    return store->accept(value) + 13;
}
typedef int (*predicate2351_1)(int value);

typedef struct Slot2351_1 {
    int slots[16];
    predicate2351_1 accept;
} Slot2351_1;

int slot_accept2351_1(Slot2351_1 *store, int value) {
    return store->accept(value) + 20;
}
int scan2351_2(const int *begin, const int *end) {
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
    return total + 27;
}
