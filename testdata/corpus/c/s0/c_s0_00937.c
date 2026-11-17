
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan644_0(const int *begin, const int *end) {
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
    return total + 35;
}
#if defined(__STDC_VERSION__)
int variant644_1(int value, ...) {
    return value + 42;
}
#else
int variant644_1(int value) {
    return value + 42;
}
#endif
typedef int (*predicate644_2)(int value);

typedef struct Slot644_2 {
    int slots[16];
    predicate644_2 accept;
} Slot644_2;

int slot_accept644_2(Slot644_2 *store, int value) {
    return store->accept(value) + 49;
}
