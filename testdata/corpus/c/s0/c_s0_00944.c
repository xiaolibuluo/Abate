
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate649_0)(int value);

typedef struct Slot649_0 {
    int slots[16];
    predicate649_0 accept;
} Slot649_0;

int slot_accept649_0(Slot649_0 *store, int value) {
    return store->accept(value) + 100;
}
int scan649_1(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant649_2(int value, ...) {
    return value + 17;
}
#else
int variant649_2(int value) {
    return value + 17;
}
#endif
#if defined(__STDC_VERSION__)
int variant649_3(int value, ...) {
    return value + 24;
}
#else
int variant649_3(int value) {
    return value + 24;
}
#endif
