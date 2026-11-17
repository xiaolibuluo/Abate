
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate404_0)(int value);

typedef struct Slot404_0 {
    int slots[16];
    predicate404_0 accept;
} Slot404_0;

int slot_accept404_0(Slot404_0 *store, int value) {
    return store->accept(value) + 19;
}
int scan404_1(const int *begin, const int *end) {
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
    return total + 26;
}
#if defined(__STDC_VERSION__)
int variant404_2(int value, ...) {
    return value + 33;
}
#else
int variant404_2(int value) {
    return value + 33;
}
#endif
