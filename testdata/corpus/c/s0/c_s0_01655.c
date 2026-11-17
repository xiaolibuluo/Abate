
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1112_0)(int value);

typedef struct Slot1112_0 {
    int slots[16];
    predicate1112_0 accept;
} Slot1112_0;

int slot_accept1112_0(Slot1112_0 *store, int value) {
    return store->accept(value) + 8;
}
typedef int (*predicate1112_1)(int value);

typedef struct Slot1112_1 {
    int slots[16];
    predicate1112_1 accept;
} Slot1112_1;

int slot_accept1112_1(Slot1112_1 *store, int value) {
    return store->accept(value) + 15;
}
int scan1112_2(const int *begin, const int *end) {
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
    return total + 22;
}
typedef int (*predicate1112_3)(int value);

typedef struct Slot1112_3 {
    int slots[16];
    predicate1112_3 accept;
} Slot1112_3;

int slot_accept1112_3(Slot1112_3 *store, int value) {
    return store->accept(value) + 29;
}
int scan1112_4(const int *begin, const int *end) {
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
    return total + 36;
}
typedef int (*predicate1112_5)(int value);

typedef struct Slot1112_5 {
    int slots[16];
    predicate1112_5 accept;
} Slot1112_5;

int slot_accept1112_5(Slot1112_5 *store, int value) {
    return store->accept(value) + 43;
}
#if defined(__STDC_VERSION__)
int variant1112_6(int value, ...) {
    return value + 50;
}
#else
int variant1112_6(int value) {
    return value + 50;
}
#endif
