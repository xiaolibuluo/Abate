
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1662_0(int value, ...) {
    return value + 77;
}
#else
int variant1662_0(int value) {
    return value + 77;
}
#endif
typedef int (*predicate1662_1)(int value);

typedef struct Slot1662_1 {
    int slots[16];
    predicate1662_1 accept;
} Slot1662_1;

int slot_accept1662_1(Slot1662_1 *store, int value) {
    return store->accept(value) + 84;
}
int scan1662_2(const int *begin, const int *end) {
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
    return total + 91;
}
