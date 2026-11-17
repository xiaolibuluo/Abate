
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan299_0(const int *begin, const int *end) {
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
typedef int (*predicate299_1)(int value);

typedef struct Slot299_1 {
    int slots[16];
    predicate299_1 accept;
} Slot299_1;

int slot_accept299_1(Slot299_1 *store, int value) {
    return store->accept(value) + 19;
}
