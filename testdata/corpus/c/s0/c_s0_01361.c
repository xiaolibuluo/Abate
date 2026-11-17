
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan934_0(const int *begin, const int *end) {
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
    return total + 22;
}
typedef int (*predicate934_1)(int value);

typedef struct Slot934_1 {
    int slots[16];
    predicate934_1 accept;
} Slot934_1;

int slot_accept934_1(Slot934_1 *store, int value) {
    return store->accept(value) + 29;
}
