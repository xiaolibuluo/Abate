
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2070_0)(int value);

typedef struct Slot2070_0 {
    int slots[16];
    predicate2070_0 accept;
} Slot2070_0;

int slot_accept2070_0(Slot2070_0 *store, int value) {
    return store->accept(value) + 46;
}
typedef int (*predicate2070_1)(int value);

typedef struct Slot2070_1 {
    int slots[16];
    predicate2070_1 accept;
} Slot2070_1;

int slot_accept2070_1(Slot2070_1 *store, int value) {
    return store->accept(value) + 53;
}
typedef int (*predicate2070_2)(int value);

typedef struct Slot2070_2 {
    int slots[16];
    predicate2070_2 accept;
} Slot2070_2;

int slot_accept2070_2(Slot2070_2 *store, int value) {
    return store->accept(value) + 60;
}
#if defined(__STDC_VERSION__)
int variant2070_3(int value, ...) {
    return value + 67;
}
#else
int variant2070_3(int value) {
    return value + 67;
}
#endif
int scan2070_4(const int *begin, const int *end) {
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
    return total + 74;
}
