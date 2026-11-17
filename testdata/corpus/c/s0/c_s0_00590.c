
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan403_0(const int *begin, const int *end) {
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
    return total + 6;
}
#if defined(__STDC_VERSION__)
int variant403_1(int value, ...) {
    return value + 13;
}
#else
int variant403_1(int value) {
    return value + 13;
}
#endif
typedef int (*predicate403_2)(int value);

typedef struct Slot403_2 {
    int slots[16];
    predicate403_2 accept;
} Slot403_2;

int slot_accept403_2(Slot403_2 *store, int value) {
    return store->accept(value) + 20;
}
