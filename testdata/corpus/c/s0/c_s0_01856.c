
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1245_0)(int value);

typedef struct Slot1245_0 {
    int slots[16];
    predicate1245_0 accept;
} Slot1245_0;

int slot_accept1245_0(Slot1245_0 *store, int value) {
    return store->accept(value) + 88;
}
int scan1245_1(const int *begin, const int *end) {
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
    return total + 95;
}
typedef int (*predicate1245_2)(int value);

typedef struct Slot1245_2 {
    int slots[16];
    predicate1245_2 accept;
} Slot1245_2;

int slot_accept1245_2(Slot1245_2 *store, int value) {
    return store->accept(value) + 5;
}
int scan1245_3(const int *begin, const int *end) {
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
    return total + 12;
}
typedef int (*predicate1245_4)(int value);

typedef struct Slot1245_4 {
    int slots[16];
    predicate1245_4 accept;
} Slot1245_4;

int slot_accept1245_4(Slot1245_4 *store, int value) {
    return store->accept(value) + 19;
}
