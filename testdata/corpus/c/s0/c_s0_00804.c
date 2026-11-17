
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan555_0(const int *begin, const int *end) {
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
    return total + 42;
}
typedef int (*predicate555_1)(int value);

typedef struct Slot555_1 {
    int slots[16];
    predicate555_1 accept;
} Slot555_1;

int slot_accept555_1(Slot555_1 *store, int value) {
    return store->accept(value) + 49;
}
typedef int (*predicate555_2)(int value);

typedef struct Slot555_2 {
    int slots[16];
    predicate555_2 accept;
} Slot555_2;

int slot_accept555_2(Slot555_2 *store, int value) {
    return store->accept(value) + 56;
}
int scan555_3(const int *begin, const int *end) {
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
    return total + 63;
}
