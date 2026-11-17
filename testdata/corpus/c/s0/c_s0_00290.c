
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan201_0(const int *begin, const int *end) {
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
    return total + 96;
}
typedef int (*predicate201_1)(int value);

typedef struct Slot201_1 {
    int slots[16];
    predicate201_1 accept;
} Slot201_1;

int slot_accept201_1(Slot201_1 *store, int value) {
    return store->accept(value) + 6;
}
