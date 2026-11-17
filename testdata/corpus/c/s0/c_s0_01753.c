
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1178_0(int value, ...) {
    return value + 90;
}
#else
int variant1178_0(int value) {
    return value + 90;
}
#endif
int scan1178_1(const int *begin, const int *end) {
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
    return total + 97;
}
typedef int (*predicate1178_2)(int value);

typedef struct Slot1178_2 {
    int slots[16];
    predicate1178_2 accept;
} Slot1178_2;

int slot_accept1178_2(Slot1178_2 *store, int value) {
    return store->accept(value) + 7;
}
