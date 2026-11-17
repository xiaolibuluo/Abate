
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant514_0(int value, ...) {
    return value + 91;
}
#else
int variant514_0(int value) {
    return value + 91;
}
#endif
int scan514_1(const int *begin, const int *end) {
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
    return total + 98;
}
typedef int (*predicate514_2)(int value);

typedef struct Slot514_2 {
    int slots[16];
    predicate514_2 accept;
} Slot514_2;

int slot_accept514_2(Slot514_2 *store, int value) {
    return store->accept(value) + 8;
}
int scan514_3(const int *begin, const int *end) {
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
    return total + 15;
}
typedef int (*predicate514_4)(int value);

typedef struct Slot514_4 {
    int slots[16];
    predicate514_4 accept;
} Slot514_4;

int slot_accept514_4(Slot514_4 *store, int value) {
    return store->accept(value) + 22;
}
