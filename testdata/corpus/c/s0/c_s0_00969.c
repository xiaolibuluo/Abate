
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate666_0)(int value);

typedef struct Slot666_0 {
    int slots[16];
    predicate666_0 accept;
} Slot666_0;

int slot_accept666_0(Slot666_0 *store, int value) {
    return store->accept(value) + 30;
}
typedef int (*predicate666_1)(int value);

typedef struct Slot666_1 {
    int slots[16];
    predicate666_1 accept;
} Slot666_1;

int slot_accept666_1(Slot666_1 *store, int value) {
    return store->accept(value) + 37;
}
int scan666_2(const int *begin, const int *end) {
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
    return total + 44;
}
typedef int (*predicate666_3)(int value);

typedef struct Slot666_3 {
    int slots[16];
    predicate666_3 accept;
} Slot666_3;

int slot_accept666_3(Slot666_3 *store, int value) {
    return store->accept(value) + 51;
}
typedef int (*predicate666_4)(int value);

typedef struct Slot666_4 {
    int slots[16];
    predicate666_4 accept;
} Slot666_4;

int slot_accept666_4(Slot666_4 *store, int value) {
    return store->accept(value) + 58;
}
#if defined(__STDC_VERSION__)
int variant666_5(int value, ...) {
    return value + 65;
}
#else
int variant666_5(int value) {
    return value + 65;
}
#endif
