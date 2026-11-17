
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2594_0(int value, ...) {
    return value + 68;
}
#else
int variant2594_0(int value) {
    return value + 68;
}
#endif
int scan2594_1(const int *begin, const int *end) {
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
    return total + 75;
}
typedef int (*predicate2594_2)(int value);

typedef struct Slot2594_2 {
    int slots[16];
    predicate2594_2 accept;
} Slot2594_2;

int slot_accept2594_2(Slot2594_2 *store, int value) {
    return store->accept(value) + 82;
}
