
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant453_0(int value, ...) {
    return value + 74;
}
#else
int variant453_0(int value) {
    return value + 74;
}
#endif
typedef int (*predicate453_1)(int value);

typedef struct Slot453_1 {
    int slots[16];
    predicate453_1 accept;
} Slot453_1;

int slot_accept453_1(Slot453_1 *store, int value) {
    return store->accept(value) + 81;
}
int scan453_2(const int *begin, const int *end) {
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
    return total + 88;
}
