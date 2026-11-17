
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan746_0(const int *begin, const int *end) {
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
    return total + 100;
}
typedef int (*predicate746_1)(int value);

typedef struct Slot746_1 {
    int slots[16];
    predicate746_1 accept;
} Slot746_1;

int slot_accept746_1(Slot746_1 *store, int value) {
    return store->accept(value) + 10;
}
typedef int (*predicate746_2)(int value);

typedef struct Slot746_2 {
    int slots[16];
    predicate746_2 accept;
} Slot746_2;

int slot_accept746_2(Slot746_2 *store, int value) {
    return store->accept(value) + 17;
}
