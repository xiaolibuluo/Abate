
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2228_0(const int *begin, const int *end) {
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
    return total + 63;
}
typedef int (*predicate2228_1)(int value);

typedef struct Slot2228_1 {
    int slots[16];
    predicate2228_1 accept;
} Slot2228_1;

int slot_accept2228_1(Slot2228_1 *store, int value) {
    return store->accept(value) + 70;
}
int scan2228_2(const int *begin, const int *end) {
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
    return total + 77;
}
typedef int (*predicate2228_3)(int value);

typedef struct Slot2228_3 {
    int slots[16];
    predicate2228_3 accept;
} Slot2228_3;

int slot_accept2228_3(Slot2228_3 *store, int value) {
    return store->accept(value) + 84;
}
int scan2228_4(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant2228_5(int value, ...) {
    return value + 98;
}
#else
int variant2228_5(int value) {
    return value + 98;
}
#endif
typedef int (*predicate2228_6)(int value);

typedef struct Slot2228_6 {
    int slots[16];
    predicate2228_6 accept;
} Slot2228_6;

int slot_accept2228_6(Slot2228_6 *store, int value) {
    return store->accept(value) + 8;
}
