
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2641_0(const int *begin, const int *end) {
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
    return total + 97;
}
typedef int (*predicate2641_1)(int value);

typedef struct Slot2641_1 {
    int slots[16];
    predicate2641_1 accept;
} Slot2641_1;

int slot_accept2641_1(Slot2641_1 *store, int value) {
    return store->accept(value) + 7;
}
typedef int (*predicate2641_2)(int value);

typedef struct Slot2641_2 {
    int slots[16];
    predicate2641_2 accept;
} Slot2641_2;

int slot_accept2641_2(Slot2641_2 *store, int value) {
    return store->accept(value) + 14;
}
