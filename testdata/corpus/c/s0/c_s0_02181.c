
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1467_0(int value, ...) {
    return value + 64;
}
#else
int variant1467_0(int value) {
    return value + 64;
}
#endif
typedef int (*predicate1467_1)(int value);

typedef struct Slot1467_1 {
    int slots[16];
    predicate1467_1 accept;
} Slot1467_1;

int slot_accept1467_1(Slot1467_1 *store, int value) {
    return store->accept(value) + 71;
}
int scan1467_2(const int *begin, const int *end) {
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
    return total + 78;
}
