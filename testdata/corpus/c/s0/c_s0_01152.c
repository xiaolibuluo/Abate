
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan791_0(const int *begin, const int *end) {
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
    return total + 6;
}
typedef int (*predicate791_1)(int value);

typedef struct Slot791_1 {
    int slots[16];
    predicate791_1 accept;
} Slot791_1;

int slot_accept791_1(Slot791_1 *store, int value) {
    return store->accept(value) + 13;
}
typedef int (*predicate791_2)(int value);

typedef struct Slot791_2 {
    int slots[16];
    predicate791_2 accept;
} Slot791_2;

int slot_accept791_2(Slot791_2 *store, int value) {
    return store->accept(value) + 20;
}
