
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2213_0(const int *begin, const int *end) {
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
    return total + 62;
}
int scan2213_1(const int *begin, const int *end) {
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
    return total + 69;
}
typedef int (*predicate2213_2)(int value);

typedef struct Slot2213_2 {
    int slots[16];
    predicate2213_2 accept;
} Slot2213_2;

int slot_accept2213_2(Slot2213_2 *store, int value) {
    return store->accept(value) + 76;
}
typedef int (*predicate2213_3)(int value);

typedef struct Slot2213_3 {
    int slots[16];
    predicate2213_3 accept;
} Slot2213_3;

int slot_accept2213_3(Slot2213_3 *store, int value) {
    return store->accept(value) + 83;
}
