
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan793_0(const int *begin, const int *end) {
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
    return total + 32;
}
typedef int (*predicate793_1)(int value);

typedef struct Slot793_1 {
    int slots[16];
    predicate793_1 accept;
} Slot793_1;

int slot_accept793_1(Slot793_1 *store, int value) {
    return store->accept(value) + 39;
}
typedef int (*predicate793_2)(int value);

typedef struct Slot793_2 {
    int slots[16];
    predicate793_2 accept;
} Slot793_2;

int slot_accept793_2(Slot793_2 *store, int value) {
    return store->accept(value) + 46;
}
