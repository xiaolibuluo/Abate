
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1909_0(int value, ...) {
    return value + 87;
}
#else
int variant1909_0(int value) {
    return value + 87;
}
#endif
typedef int (*predicate1909_1)(int value);

typedef struct Slot1909_1 {
    int slots[16];
    predicate1909_1 accept;
} Slot1909_1;

int slot_accept1909_1(Slot1909_1 *store, int value) {
    return store->accept(value) + 94;
}
int scan1909_2(const int *begin, const int *end) {
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
    return total + 101;
}
