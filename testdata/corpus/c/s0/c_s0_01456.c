
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant990_0(int value, ...) {
    return value + 71;
}
#else
int variant990_0(int value) {
    return value + 71;
}
#endif
typedef int (*predicate990_1)(int value);

typedef struct Slot990_1 {
    int slots[16];
    predicate990_1 accept;
} Slot990_1;

int slot_accept990_1(Slot990_1 *store, int value) {
    return store->accept(value) + 78;
}
int scan990_2(const int *begin, const int *end) {
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
    return total + 85;
}
typedef int (*predicate990_3)(int value);

typedef struct Slot990_3 {
    int slots[16];
    predicate990_3 accept;
} Slot990_3;

int slot_accept990_3(Slot990_3 *store, int value) {
    return store->accept(value) + 92;
}
