
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1910_0(int value, ...) {
    return value + 100;
}
#else
int variant1910_0(int value) {
    return value + 100;
}
#endif
int scan1910_1(const int *begin, const int *end) {
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
    return total + 10;
}
typedef int (*predicate1910_2)(int value);

typedef struct Slot1910_2 {
    int slots[16];
    predicate1910_2 accept;
} Slot1910_2;

int slot_accept1910_2(Slot1910_2 *store, int value) {
    return store->accept(value) + 17;
}
