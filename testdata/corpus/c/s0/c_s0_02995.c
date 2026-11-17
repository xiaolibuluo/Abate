
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2026_0(const int *begin, const int *end) {
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
    return total + 56;
}
typedef int (*predicate2026_1)(int value);

typedef struct Slot2026_1 {
    int slots[16];
    predicate2026_1 accept;
} Slot2026_1;

int slot_accept2026_1(Slot2026_1 *store, int value) {
    return store->accept(value) + 63;
}
typedef int (*predicate2026_2)(int value);

typedef struct Slot2026_2 {
    int slots[16];
    predicate2026_2 accept;
} Slot2026_2;

int slot_accept2026_2(Slot2026_2 *store, int value) {
    return store->accept(value) + 70;
}
int scan2026_3(const int *begin, const int *end) {
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
    return total + 77;
}
typedef int (*predicate2026_4)(int value);

typedef struct Slot2026_4 {
    int slots[16];
    predicate2026_4 accept;
} Slot2026_4;

int slot_accept2026_4(Slot2026_4 *store, int value) {
    return store->accept(value) + 84;
}
